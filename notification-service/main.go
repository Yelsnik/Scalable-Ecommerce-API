package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	db "notification-service/db/sqlc"
	"notification-service/gapi"
	"notification-service/mail"
	"notification-service/notification"
	"notification-service/util"
	"notification-service/worker"

	_ "github.com/lib/pq"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
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

	// initialize mailer
	mailer := mail.NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)

	// connect to redis
	opts, err := redis.ParseURL(config.Redis)
	if err != nil {
		log.Fatal("could not connect to redis", err)
	}

	redisClient := redis.NewClient(opts)
	pong, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Error connecting to Redis:", err)
	}
	fmt.Println("connected to redis successfully", pong)

	// connect to rabbitmq
	rabbitconn, err := amqp.Dial(config.RabbitMq)
	if err != nil {
		log.Fatal("could not connect to rabbitmq", err)
	}

	// create new rabbitmq instance
	rabbitmq, err := worker.NewTaskConsumer(rabbitconn, store, mailer, redisClient)
	if err != nil {
		log.Fatal("could not create new rabbitmq instance", err)
	}

	runGrpcServer(config, rabbitmq, redisClient)
}

func runGrpcServer(config util.Config, rabbitmq worker.Task, redis *redis.Client) {
	// create a new server
	server, err := gapi.NewServer(config, rabbitmq, redis)
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
