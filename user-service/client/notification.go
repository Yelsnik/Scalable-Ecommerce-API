package client

import (
	"context"
	"user-service/notification/notification"
)

func (client *Client) SendEmail(ctx context.Context, in *notification.SendEmailRequest) (*notification.SendEmailResponse, error) {
	return client.NotifClient.SendEmail(ctx, in)
}
