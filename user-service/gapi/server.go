package gapi

import (
	"fmt"

	db "user-service/db/sqlc"

	"user-service/pb"
	"user-service/token"
	"user-service/util"
)

// serves gRPC requests for our banking service
type Server struct {
	pb.UnimplementedUserServiceServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// creates a new gRPC server
func NewServer(config util.Config, store db.Store) (*Server, error) {

	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
