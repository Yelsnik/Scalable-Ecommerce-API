package gapi

import (
	"auction-service/auction"
	db "auction-service/db/sqlc"
	"auction-service/util"
	"auction-service/worker"

	"github.com/redis/go-redis/v9"
)

// serves gRPC requests for our banking service
type Server struct {
	auction.UnimplementedAuctionServiceServer
	config   util.Config
	store    db.Store
	redis    *redis.Client
	rabbitmq worker.Amqp
}

// creates a new gRPC server
func NewServer(config util.Config, store db.Store, redis *redis.Client, rabbitmq worker.Amqp) (*Server, error) {

	server := &Server{
		config:   config,
		store:    store,
		redis:    redis,
		rabbitmq: rabbitmq,
	}

	return server, nil
}
