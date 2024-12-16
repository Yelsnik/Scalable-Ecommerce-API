package main

import (
	"context"
	"database/sql"
	"log"
	"net"
	"net/http"

	db "user-service/db/sqlc"
	"user-service/gapi"
	"user-service/pb"
	"user-service/util"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
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

	store := db.NewStore(conn)
	go runGatewayServer(config, store)
	runGrpcServer(config, store)
}

func runGrpcServer(config util.Config, store db.Store) {
	// create a new server
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	// create a new grpc server
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, server)
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

func runGatewayServer(config util.Config, store db.Store) {
	// create a new server
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	//
	grpcMux := runtime.NewServeMux()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = pb.RegisterUserServiceHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal("cannot register handler server:", err)
	}

	// create http serve mux to receive requests
	// from clients
	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	/*
		statikFs, err := fs.New()
		if err != nil {
			log.Fatal("cannot create statik fs", err)
		}

		swaggerHandler := http.StripPrefix("/swagger/", http.FileServer(statikFs))
		mux.Handle("/swagger/", swaggerHandler)
	*/

	// start the server to listen to grpc
	// requests on a specific port
	listener, err := net.Listen("tcp", config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot create listener:", err)
	}

	log.Printf("starting HTTP Gateway server at %s ...", listener.Addr().String())

	// start the server
	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatal("cannot start HTTP Gateway server:", err)
	}

}
