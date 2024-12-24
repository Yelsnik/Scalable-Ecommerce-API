package db

import (
	"cart-service/util"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createNewCart(t *testing.T) Cart {
	args := CreateCartParams{
		UserID:     util.Test(),
		TotalPrice: 0,
	}

	cart, err := testStore.CreateCart(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, cart)
	require.Equal(t, args.UserID, cart.UserID)
	require.Equal(t, args.TotalPrice, cart.TotalPrice)

	return cart
}

func TestCreateCart(t *testing.T) {
	createNewCart(t)
}
