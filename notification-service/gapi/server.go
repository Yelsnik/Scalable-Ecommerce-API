package gapi

import (
	"notification-service/client"
	"notification-service/notification/notification"
	"notification-service/util"
	"notification-service/worker"
)

// serves gRPC requests for our banking service
type Server struct {
	notification.UnimplementedNotificationServiceServer
	config   util.Config
	rabbitmq worker.Task
	client   *client.Client
}

// creates a new gRPC server
func NewServer(config util.Config, rabbitmq worker.Task, client *client.Client) (*Server, error) {

	server := &Server{
		config:   config,
		rabbitmq: rabbitmq,
		client:   client,
	}

	return server, nil
}
