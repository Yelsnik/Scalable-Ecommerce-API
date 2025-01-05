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
