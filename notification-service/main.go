package main

import (
	"database/sql"
	"log"
	"net"
	"notification-service/client"
	db "notification-service/db/sqlc"
	"notification-service/gapi"
	"notification-service/mail"
	"notification-service/notification/notification"
	"notification-service/util"
	"notification-service/worker"

	_ "github.com/lib/pq"
	amqp "github.com/rabbitmq/amqp091-go"
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

	store := db.NewStore(conn)

	// initialize mailer
	mailer := mail.NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)

	// connect to rabbitmq
	rabbitconn, err := amqp.Dial(config.RabbitMq)
	if err != nil {
		log.Fatal("could not connect to rabbitmq", err)
	}

	// create new rabbitmq instance
	rabbitmq, err := worker.NewTaskConsumer(rabbitconn, store, mailer)
	if err != nil {
		log.Fatal("could not create new rabbitmq instance", err)
	}

	// create client grpc conn
	authConn, err := grpc.NewClient("0.0.0.0:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to auth service: %v", err)
	}

	// Initialize the client
	client := client.NewClient(authConn)

	runGrpcServer(config, rabbitmq, client)
}

func runGrpcServer(config util.Config, rabbitmq worker.Task, client *client.Client) {
	// create a new server
	server, err := gapi.NewServer(config, rabbitmq, client)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	// create a new grpc server
	grpcServer := grpc.NewServer()
	notification.RegisterNotificationServiceServer(grpcServer, server)
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
