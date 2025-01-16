package gapi

import (
	"fmt"

	"user-service/client"
	db "user-service/db/sqlc"
	"user-service/worker"

	"user-service/pb"
	"user-service/token"
	"user-service/util"
)

// serves gRPC requests for our banking service
type Server struct {
	pb.UnimplementedUserServiceServer
	pb.UnimplementedAuthServiceServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	rabbitmq   worker.Amqp
	client     *client.Client
}

// creates a new gRPC server
func NewServer(config util.Config, store db.Store, rabbitmq worker.Amqp, client *client.Client) (*Server, error) {

	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
		rabbitmq:   rabbitmq,
		client:     client,
	}

	return server, nil
}
