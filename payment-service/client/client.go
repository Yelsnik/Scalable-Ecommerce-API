package client

import (
	"payment-service/cart/cart-service"
	"payment-service/product/product-service"
	pb "payment-service/user/user-service"

	"google.golang.org/grpc"
)

// Client wraps all gRPC service clients.
type Client struct {
	ProductClient product.ProductServiceClient
	ShopClient    product.ShopServiceClient
	AuthClient    pb.AuthServiceClient
	CartClient    cart.CartItemServiceClient
}

// NewClient creates a new Client instance with connections to all gRPC services.
func NewClient(productConn, shopConn, authConn, cartConn *grpc.ClientConn) *Client {
	return &Client{
		ProductClient: product.NewProductServiceClient(productConn),
		ShopClient:    product.NewShopServiceClient(shopConn),
		AuthClient:    pb.NewAuthServiceClient(authConn),
		CartClient:    cart.NewCartItemServiceClient(cartConn),
	}
}
