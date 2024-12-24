package db

import (
	"cart-service/util"
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type Product struct {
	id       string
	price    float64
	currency string
}

func createNewProduct() Product {
	product := Product{
		id:       util.RandomString(8),
		price:    float64(util.RandomMoney()),
		currency: util.RandomCurrency(),
	}

	return product
}

func TestAddToCartTx(t *testing.T) {
	var product []Product
	for i := 0; i < 3; i++ {
		product = append(product, createNewProduct())
	}

	cart := createNewCart(t)

	n := 3

	for i := 0; i < len(product); i++ {

		q := util.RandomInt(1, 100)

		subTotal := float64(q) * product[i].price
		arg := CreateCartitemParams{
			Cart:     cart.ID,
			Product:  product[i].id,
			Quantity: q,
			Price:    product[i].price,
			Currency: product[i].currency,
			SubTotal: subTotal,
		}

		errs := make(chan error, 5)
		results := make(chan CartTxResult, 5)

		for i := 0; i < n; i++ {
			txName := fmt.Sprintf("tx %d", i+1)

			go func() {
				ctx := context.WithValue(context.Background(), txKey, txName)
				result, err := testStore.AddToCartTx(ctx, arg)

				results <- result
				errs <- err
			}()
		}

		// check result
		for i := 0; i < n; i++ {
			err := <-errs
			require.NoError(t, err)

			result := <-results
			require.NotEmpty(t, result)

			cartResult := result.Cart
			require.NotEmpty(t, cartResult)
			require.Equal(t, cart.ID, cartResult.ID)

			cartItemResult := result.CartItem
			require.NotEmpty(t, cartItemResult)

		}
	}

}

func TestUpdateCartTx(t *testing.T) {

	n := 20

	for i := 0; i < n; i++ {

		cartItem := createNewCartItem(t)
		q := util.RandomInt(1, 100)
		subTotal := cartItem.Price * float64(q)
		arg := UpdateCartitemParams{
			ID:       cartItem.ID,
			Quantity: q,
			SubTotal: subTotal,
		}

		errs := make(chan error, 5)
		results := make(chan CartTxResult, 5)

		txName := fmt.Sprintf("tx %d", i+1)
		go func() {
			ctx := context.WithValue(context.Background(), txKey, txName)
			result, err := testStore.UpdateCartTx(ctx, cartItem.ID, arg)

			results <- result
			errs <- err
		}()

		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		cartResult := result.Cart
		require.NotEmpty(t, cartResult)

		cartItemResult := result.CartItem
		require.NotEmpty(t, cartItemResult)
		require.Equal(t, cartItem.Cart, cartItemResult.Cart)
		require.Equal(t, cartItem.Product, cartItemResult.Product)
		require.Equal(t, arg.Quantity, cartItemResult.Quantity)
		require.Equal(t, cartItem.Price, cartItemResult.Price)
		require.Equal(t, arg.SubTotal, cartItemResult.SubTotal)
		require.NotZero(t, cartItemResult.CreatedAt)
	}
}
