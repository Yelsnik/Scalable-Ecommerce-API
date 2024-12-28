package client

import (
	"cart-service/product"
	"context"
)

func (c *Client) GetProductByID(ctx context.Context, req string) (*product.ProductResponse, error) {
	in := &product.GetProductByIdRequest{
		Id: req,
	}

	return c.client.GetProductByID(ctx, in)
}
