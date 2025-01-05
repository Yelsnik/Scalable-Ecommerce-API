package client

import (
	"context"
	"payment-service/product/product-service"
)

func (client *Client) GetShopByID(ctx context.Context, req string) (*product.ShopResponse, error) {

	in := &product.GetShopByIdRequest{
		Id: req,
	}
	return client.ShopClient.GetShopByID(ctx, in)
}
