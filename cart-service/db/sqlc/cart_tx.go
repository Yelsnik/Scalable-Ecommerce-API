package db

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type CartTxResult struct {
	CartItem Cartitem `json:"cart_item"`
	Cart     Cart     `json:"cart"`
}

var txKey = struct{}{}

func (store *SQLStore) AddToCartTx(ctx context.Context, arg CreateCartitemParams) (CartTxResult, error) {
	var result CartTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		txName := ctx.Value(txKey)

		// create new cart items
		fmt.Println(txName, "create cart items")
		result.CartItem, err = q.CreateCartitem(ctx, CreateCartitemParams{
			Cart:     arg.Cart,
			Product:  arg.Product,
			Quantity: arg.Quantity,
			Price:    arg.Price,
			Currency: arg.Currency,
			SubTotal: arg.SubTotal,
		})
		if err != nil {
			return err
		}

		// add subtotal price
		total, err := q.AddSubtotalPrice(ctx, result.CartItem.Cart)
		if err != nil {
			return err
		}

		// update the carts with the added total
		fmt.Println(txName, "update carts")
		result.Cart, err = q.UpdateCart(ctx, UpdateCartParams{
			ID:         result.CartItem.Cart,
			TotalPrice: total,
		})
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}

func (store *SQLStore) UpdateCartTx(ctx context.Context, cartItemID uuid.UUID, arg UpdateCartitemParams) (CartTxResult, error) {
	var result CartTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		txName := ctx.Value(txKey)

		// get cart item for update
		fmt.Println(txName, "Get cart items")
		result.CartItem, err = q.GetCartitemForUpdate(ctx, cartItemID)
		if err != nil {
			return err
		}

		// update cart item
		fmt.Println(txName, "update cart items")
		result.CartItem, err = q.UpdateCartitem(ctx, UpdateCartitemParams{
			ID:       result.CartItem.ID,
			Quantity: arg.Quantity,
			SubTotal: arg.SubTotal,
		})
		if err != nil {
			return err
		}

		// calculate total price of cart
		fmt.Println(txName, "calculate subtotal of cart items")
		total, err := q.AddSubtotalPrice(ctx, result.CartItem.Cart)
		if err != nil {
			return err
		}

		// update carts
		fmt.Println(txName, "update carts total")
		result.Cart, err = q.UpdateCart(ctx, UpdateCartParams{
			ID:         result.CartItem.Cart,
			TotalPrice: total,
		})
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}

type RemoveCartTxResult struct {
	Cart Cart `json:"cart"`
}

func (store *SQLStore) RemoveCartTx(ctx context.Context, cartItemID uuid.UUID) (RemoveCartTxResult, error) {
	var result RemoveCartTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		var cartId uuid.UUID

		// get cart item
		cartItem, err := q.GetCartitemForUpdate(ctx, cartItemID)
		if err != nil {
			return err
		}
		cartId = cartItem.Cart

		// delete cart item
		err = q.DeleteCartitem(ctx, cartItem.ID)
		if err != nil {
			return err
		}

		// add cart item subtotal to get total price
		total, err := q.AddSubtotalPrice(ctx, cartId)
		if err != nil {
			return err
		}

		// update cart total
		result.Cart, err = q.UpdateCart(ctx, UpdateCartParams{
			ID:         cartId,
			TotalPrice: total,
		})

		if err != nil {
			return err
		}

		return err
	})

	return result, err
}
