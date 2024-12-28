package client

import (
	"cart-service/product/product"

	"google.golang.org/grpc"
)

type Client struct {
	client product.ProductServiceClient
}

func NewClient(conn *grpc.ClientConn) *Client {
	return &Client{
		client: product.NewProductServiceClient(conn),
	}
}
