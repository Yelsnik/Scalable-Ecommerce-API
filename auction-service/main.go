package main

import (
	"auction-service/auction"
	db "auction-service/db/sqlc"
	"auction-service/gapi"
	"auction-service/util"
	"auction-service/worker"
	"context"
	"fmt"
	"log"
	"net"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// load config
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// connect to database
	conn, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("could not connect", err)
	}

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
	rabbitmq, err := worker.NewAmqpTask(rabbitconn, worker.QueueBid, true)
	if err != nil {
		log.Fatal("could not create new rabbitmq instance", err)
	}

	store := db.NewStore(conn)

	runGrpcServer(config, store, redisClient, rabbitmq)
}

func runGrpcServer(config util.Config, store db.Store, redis *redis.Client, rabbitmq worker.Amqp) {
	// create a new server
	server, err := gapi.NewServer(config, store, redis, rabbitmq)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	// create a new grpc server
	grpcServer := grpc.NewServer()
	auction.RegisterAuctionServiceServer(grpcServer, server)
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
