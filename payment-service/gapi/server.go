package gapi

import (
	"payment-service/client"
	db "payment-service/db/sqlc"
	"payment-service/payment/payment-service"

	"payment-service/util"
)

// serves gRPC requests for our banking service
type Server struct {
	payment.UnimplementedPaymentServiceServer
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
