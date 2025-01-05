package client

import (
	"context"
	"payment-service/product/product-service"
)

func (c *Client) GetProductByID(ctx context.Context, req string) (*product.ProductResponse, error) {
	in := &product.GetProductByIdRequest{
		Id: req,
	}

	return c.ProductClient.GetProductByID(ctx, in)
}
