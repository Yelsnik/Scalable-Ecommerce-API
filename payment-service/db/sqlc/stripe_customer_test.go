package db

import (
	"context"
	"payment-service/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createnewStripeCustomer(t *testing.T) StripeCustomer {
	userId := util.Test()

	args := CreateStripeCustomerParams{
		ID:     util.RandomString(7),
		UserID: userId,
	}

	stripeCustomer, err := testStore.CreateStripeCustomer(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, stripeCustomer)

	return stripeCustomer
}

func TestCreateStripeC(t *testing.T) {
	createnewStripeCustomer(t)
}
