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

func (client *Client) GetUserByEmail(ctx context.Context, email string) (*pb.GetUserByEmailResponse, error) {
	in := &pb.GetUserByEmailRequest{
		Email: email,
	}

	return client.AuthClient.GetUserByEmail(ctx, in)
}
