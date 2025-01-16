package gapi

import (
	"context"
	db "notification-service/db/sqlc"
	"notification-service/notification"
	"notification-service/util"
	"notification-service/worker"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) SendEmail(ctx context.Context, req *notification.SendEmailRequest) (*notification.SendEmailResponse, error) {

	id, err := util.ConvertStringToUUID(req.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid id %s", err)
	}

	arg := worker.ConsumeTask{
		Ctx: ctx,
		Task: worker.TaskSendVerifyEmailRequest{
			Arg: db.CreateVerifyEmailParams{
				UserID:     id,
				Email:      req.GetEmail(),
				UserName:   req.GetUserName(),
				SecretCode: util.RandomString(32),
			},
		},
	}

	err = server.rabbitmq.Consume("notification-service", arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to send email %s", err)
	}

	response := &notification.SendEmailResponse{
		Message: "succesfully sent email",
	}

	return response, nil
}
