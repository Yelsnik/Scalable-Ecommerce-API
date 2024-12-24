package db

import (
	"cart-service/util"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createNewCartItem(t *testing.T) Cartitem {
	cart := createNewCart(t)

	product := struct {
		id       string
		price    float64
		currency string
	}{
		id:       util.RandomString(8),
		price:    float64(util.RandomMoney()),
		currency: util.RandomCurrency(),
	}

	q := util.RandomInt(1, 100)

	subTotal := product.price * float64(q)

	arg := CreateCartitemParams{
		Cart:     cart.ID,
		Product:  product.id,
		Quantity: q,
		Price:    product.price,
		Currency: product.currency,
		SubTotal: subTotal,
	}

	cartItem, err := testStore.CreateCartitem(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, cartItem)

	require.Equal(t, arg.Cart, cartItem.Cart)
	require.Equal(t, arg.Product, cartItem.Product)
	require.Equal(t, arg.Quantity, cartItem.Quantity)
	require.Equal(t, arg.Price, cartItem.Price)
	require.Equal(t, arg.SubTotal, cartItem.SubTotal)
	require.NotEmpty(t, cartItem.CreatedAt)
	require.NotZero(t, cartItem.CreatedAt)

	return cartItem
}

func TestCreateCartItem(t *testing.T) {
	createNewCartItem(t)
}
