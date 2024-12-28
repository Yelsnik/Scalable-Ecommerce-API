package gapi

import (
	db "cart-service/db/sqlc"

	pb "cart-service/cart"

	"cart-service/client"
	"cart-service/util"
)

// serves gRPC requests for our banking service
type Server struct {
	pb.UnimplementedCartItemServiceServer
	pb.UnimplementedCartServiceServer
	config util.Config
	store  db.Store
	client *client.Client
}

// creates a new gRPC server
func NewServer(config util.Config, store db.Store, client *client.Client) (*Server, error) {

	server := &Server{
		config: config,
		store:  store,
		client: client,
	}

	return server, nil
}
