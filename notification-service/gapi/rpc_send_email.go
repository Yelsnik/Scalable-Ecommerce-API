package gapi

import (
	"context"

	"notification-service/notification"
	"notification-service/worker"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) SendEmail(ctx context.Context, req *notification.SendEmailRequest) (*notification.SendEmailResponse, error) {

	arg := worker.ConsumeTask{
		Ctx: ctx,
	}

	err := server.rabbitmq.Consume("email-service", 10000, true, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to send email %s", err)
	}

	response := &notification.SendEmailResponse{
		Message: "succesfully sent email",
	}

	return response, nil
}
