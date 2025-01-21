package gapi

import (
	"auction-service/auction"
	db "auction-service/db/sqlc"
	"auction-service/util"
)

// serves gRPC requests for our banking service
type Server struct {
	auction.UnimplementedAuctionServiceServer
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
