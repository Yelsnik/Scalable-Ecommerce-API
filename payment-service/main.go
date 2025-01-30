package main

import (
	"database/sql"
	"log"
	"net"

	"payment-service/client"
	db "payment-service/db/sqlc"
	"payment-service/gapi"
	"payment-service/gapi/helpers"
	"payment-service/payment/payment-service"
	"payment-service/stripe"
	"payment-service/util"

	_ "github.com/lib/pq"
	str "github.com/stripe/stripe-go/v81"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	str.Key = config.StripeSecretKey

	// connect to database
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("could not connect", err)
	}

	// create grpc conn
	productConn, err := grpc.NewClient("0.0.0.0:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to product service: %v", err)
	}
	defer productConn.Close()

	shopConn, err := grpc.NewClient("0.0.0.0:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to product service: %v", err)
	}
	defer shopConn.Close()

	authConn, err := grpc.NewClient("0.0.0.0:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to auth service: %v", err)
	}
	defer authConn.Close()

	cartConn, err := grpc.NewClient("0.0.0.0:7070", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to cart service: %v", err)
	}
	defer cartConn.Close()

	// Initialize the client
	client := client.NewClient(productConn, shopConn, authConn, cartConn)

	// initialize and inject helper functions for the server
	helpers := helpers.NewHelpers()

	// initialize and inject stripe interface
	stripe := stripe.NewStripeClient()

	store := db.NewStore(conn)

	grpcServerParams := gapi.NewGrpcServerParams{
		Config:  config,
		Store:   store,
		Client:  client,
		Helpers: helpers,
		Stripe:  stripe,
	}

	runGrpcServer(grpcServerParams)
}

func runGrpcServer(params gapi.NewGrpcServerParams) {
	// create a new server
	server, err := gapi.NewServer(params)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	// create a new grpc server
	grpcServer := grpc.NewServer()
	payment.RegisterPaymentServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	// start the server to listen to grpc
	// requests on a specific port
	listener, err := net.Listen("tcp", params.Config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create listener:", err)
	}

	log.Printf("starting gRPC server at %s ...", listener.Addr().String())

	// start the server
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start grpc server:", err)
	}

}
