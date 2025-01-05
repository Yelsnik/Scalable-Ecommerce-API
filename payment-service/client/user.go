package client

import (
	"context"
	pb "payment-service/user/user-service"
)

func (client *Client) GetStripeSellerAccount(ctx context.Context, req string) (*pb.StripeSellerAccountResponse, error) {

	in := &pb.StripeSellerAccountRequest{
		UserId: req,
	}

	return client.AuthClient.GetStripeSellerAccount(ctx, in)
}

func (client *Client) GetUserByID(ctx context.Context, req string) (*pb.GetUserByIdResponse, error) {
	in := &pb.GetUserByIdRequest{
		Id: req,
	}

	return client.AuthClient.GetUserByID(ctx, in)
}
