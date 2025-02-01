package gapi

import (
	"testing"

	"payment-service/client"
	db "payment-service/db/sqlc"
	"payment-service/gapi/helpers"
	"payment-service/stripe"

	"payment-service/util"

	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store, client client.ClientInterface, helpers helpers.Helper, stripe stripe.Stripe) *Server {
	config := util.Config{
		WebhookSigningKey: "whsec_1d6887b20f7014249116c00aa0ddd2fe96e379f3b568a202d9ce60bdbe09dda8",
	}

	params := NewGrpcServerParams{
		Config:  config,
		Store:   store,
		Client:  client,
		Helpers: helpers,
		Stripe:  stripe,
	}

	server, err := NewServer(params)
	require.NoError(t, err)

	return server
}
