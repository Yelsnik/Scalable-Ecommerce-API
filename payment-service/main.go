package main

import (
	"database/sql"
	"log"
	"net"

	"payment-service/client"
	db "payment-service/db/sqlc"
	"payment-service/gapi"
	"payment-service/payment/payment-service"
	"payment-service/util"

	_ "github.com/lib/pq"
	"github.com/stripe/stripe-go/v81"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	stripe.Key = config.StripeSecretKey

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

	store := db.NewStore(conn)
	runGrpcServer(config, store, client)
}

func runGrpcServer(config util.Config, store db.Store, c *client.Client) {
	// create a new server
	server, err := gapi.NewServer(config, store, c)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	// create a new grpc server
	grpcServer := grpc.NewServer()
	payment.RegisterPaymentServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	// start the server to listen to grpc
	// requests on a specific port
	listener, err := net.Listen("tcp", config.GRPCServerAddress)
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
