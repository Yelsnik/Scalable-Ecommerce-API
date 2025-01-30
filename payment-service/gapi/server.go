package gapi

import (
	"payment-service/client"
	db "payment-service/db/sqlc"
	"payment-service/gapi/helpers"
	"payment-service/payment/payment-service"
	"payment-service/stripe"

	"payment-service/util"
)

// serves gRPC requests for our banking service
type Server struct {
	payment.UnimplementedPaymentServiceServer
	config  util.Config
	store   db.Store
	client  client.ClientInterface
	helpers helpers.Helper
	stripe  stripe.Stripe
}

// creates a new gRPC server
type NewGrpcServerParams struct {
	Config  util.Config
	Store   db.Store
	Client  client.ClientInterface
	Helpers helpers.Helper
	Stripe  stripe.Stripe
}

func NewServer(params NewGrpcServerParams) (*Server, error) {

	server := &Server{
		config:  params.Config,
		store:   params.Store,
		client:  params.Client,
		helpers: params.Helpers,
		stripe:  params.Stripe,
	}

	return server, nil
}
