package gapi

import (
	"context"
	db "user-service/db/sqlc"
	"user-service/pb"
	"user-service/util"
	"user-service/worker"

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
				Email      string
				UserName   string
				Id         string
				SecretCode string
			}{
				Email:      user.Email,
				UserName:   user.Name,
				Id:         user.ID.String(),
				SecretCode: util.RandomString(32),
			}

			err := server.rabbitmq.Publish(worker.TaskSendVerifyEmail, payload, ctx)

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
