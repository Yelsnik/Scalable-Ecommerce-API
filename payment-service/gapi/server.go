package gapi

import (
	db "payment-service/db/sqlc"
	"payment-service/payment"

	"payment-service/util"
)

// serves gRPC requests for our banking service
type Server struct {
	payment.UnimplementedPaymentServiceServer
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
