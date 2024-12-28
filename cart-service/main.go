package main

import (
	pb "cart-service/cart"
	"cart-service/client"
	db "cart-service/db/sqlc"
	"cart-service/gapi"
	"cart-service/util"
	"database/sql"
	"log"
	"net"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// connect to database
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("could not connect", err)
	}

	productConn, err := grpc.NewClient("0.0.0.0:50051", grpc.WithTransportCredentials(insecure.NewCredentials())) // Use WithTransportCredentials for secure connections
	if err != nil {
		log.Fatalf("Failed to connect to Product Service: %v", err)
	}
	defer productConn.Close()

	client := client.NewClient(productConn)

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
	pb.RegisterCartItemServiceServer(grpcServer, server)
	pb.RegisterCartServiceServer(grpcServer, server)
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
