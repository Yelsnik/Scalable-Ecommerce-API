package client

import (
	"context"
	"payment-service/cart/cart-service"
)

func (client *Client) GetCartItem(ctx context.Context, req string) (*cart.CartItemResponse, error) {
	in := &cart.GetCartItemByIDRequest{
		Id: req,
	}

	return client.CartClient.GetCartItem(ctx, in)
}

func (client *Client) RemoveCartTx(ctx context.Context, cartItemID string) (*cart.RemoveCartTxResult, error) {
	in := &cart.RemoveCartTxRequest{
		Id: cartItemID,
	}

	return client.CartClient.RemoveCartTx(ctx, in)
}
