package gapi

import (
	db "cart-service/db/sqlc"

	pb "cart-service/cart"

	"cart-service/util"
)

// serves gRPC requests for our banking service
type Server struct {
	pb.UnimplementedCartItemServiceServer
	pb.UnimplementedCartServiceServer
	config util.Config
	store  db.Store
}

// creates a new gRPC server
func NewServer(config util.Config, store db.Store) (*Server, error) {

	server := &Server{
		config: config,
		store:  store,
	}

	return server, nil
}
