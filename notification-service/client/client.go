package client

import (
	pb "notification-service/user"

	"google.golang.org/grpc"
)

// Client wraps all gRPC service clients.
type Client struct {
	AuthClient pb.AuthServiceClient
}

// NewClient creates a new Client instance with connections to all gRPC services.
func NewClient(authConn *grpc.ClientConn) *Client {
	return &Client{
		AuthClient: pb.NewAuthServiceClient(authConn),
	}
}
