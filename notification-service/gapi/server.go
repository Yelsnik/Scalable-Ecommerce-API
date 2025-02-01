package gapi

import (
	"notification-service/notification"
	"notification-service/util"
	"notification-service/worker"

	"github.com/redis/go-redis/v9"
)

// serves gRPC requests for our banking service
type Server struct {
	notification.UnimplementedNotificationServiceServer
	config   util.Config
	rabbitmq worker.Task
	redis    *redis.Client
}

// creates a new gRPC server
func NewServer(config util.Config, rabbitmq worker.Task, redis *redis.Client) (*Server, error) {

	server := &Server{
		config:   config,
		rabbitmq: rabbitmq,
		redis:    redis,
	}

	return server, nil
}
