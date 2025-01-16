package gapi

import (
	"context"
	db "user-service/db/sqlc"
	"user-service/pb"
	"user-service/util"
	"user-service/worker"

	"user-service/notification/notification"

	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	hashedPassword, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)
	}

	arg := db.CreateUserTxParams{
		CreateUserParams: db.CreateUserParams{
			Name:     req.GetName(),
			Email:    req.GetEmail(),
			Role:     req.GetRole(),
			Password: hashedPassword,
		},
		AfterCreate: func(user db.User) error {
			payload := struct {
				Email string
			}{
				Email: user.Email,
			}

			arg := &notification.SendEmailRequest{
				Email:    user.Email,
				UserName: user.Name,
				UserId:   user.ID.String(),
			}

			err := server.rabbitmq.Publish(worker.TaskSendVerifyEmail, arg, payload, ctx)

			return err
		},
	}

	result, err := server.store.CreateUserTx(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return nil, status.Errorf(codes.AlreadyExists, "user already exists: %s", err)
			}
		}
		return nil, status.Errorf(codes.Internal, "failed to create user and send email: %s", err)
	}

	rsp := &pb.CreateUserResponse{
		User: convertUser(result.User),
	}

	return rsp, nil
}
