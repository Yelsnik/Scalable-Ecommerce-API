package client

import (
	"user-service/notification/notification"

	"google.golang.org/grpc"
)

// Client wraps all gRPC service clients.
type Client struct {
	NotifClient notification.NotificationServiceClient
}

// NewClient creates a new Client instance with connections to all gRPC services.
func NewClient(notifConn *grpc.ClientConn) *Client {
	return &Client{
		NotifClient: notification.NewNotificationServiceClient(notifConn),
	}
}
