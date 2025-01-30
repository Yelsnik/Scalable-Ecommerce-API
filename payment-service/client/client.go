package client

import (
	"context"
	"payment-service/cart/cart-service"
	"payment-service/product/product-service"
	pb "payment-service/user/user-service"

	"google.golang.org/grpc"
)

type ClientInterface interface {
	GetCartItem(ctx context.Context, req string) (*cart.CartItemResponse, error)
	RemoveCartTx(ctx context.Context, cartItemID string) (*cart.RemoveCartTxResult, error)
	GetProductByID(ctx context.Context, req string) (*product.ProductResponse, error)
	GetShopByID(ctx context.Context, req string) (*product.ShopResponse, error)
	GetStripeSellerAccount(ctx context.Context, req string) (*pb.StripeSellerAccountResponse, error)
	GetUserByID(ctx context.Context, req string) (*pb.GetUserByIdResponse, error)
	GetUserByEmail(ctx context.Context, email string) (*pb.GetUserByEmailResponse, error)
}

// Client wraps all gRPC service clients.
type Client struct {
	ProductClient product.ProductServiceClient
	ShopClient    product.ShopServiceClient
	AuthClient    pb.AuthServiceClient
	CartClient    cart.CartItemServiceClient
}

// NewClient creates a new Client instance with connections to all gRPC services.
func NewClient(productConn, shopConn, authConn, cartConn *grpc.ClientConn) ClientInterface {
	return &Client{
		ProductClient: product.NewProductServiceClient(productConn),
		ShopClient:    product.NewShopServiceClient(shopConn),
		AuthClient:    pb.NewAuthServiceClient(authConn),
		CartClient:    cart.NewCartItemServiceClient(cartConn),
	}
}
