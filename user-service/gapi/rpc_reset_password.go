package gapi

import (
	"context"
	"fmt"
	db "user-service/db/sqlc"
	"user-service/pb"
	"user-service/util"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) ResetPassword(ctx context.Context, req *pb.ResetPasswordRequest) (*pb.ResetPasswordResponse, error) {
	_, err := server.tokenMaker.VerifyToken(req.ResetToken)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "token is invalid: %s", err)
	}

	passwordResetToken, err := server.store.GetPasswordResetTokenByToken(ctx, req.ResetToken)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get token: %s", err)
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)
	}

	user, err := server.store.UpdateUser(ctx, db.UpdateUserParams{
		ID:       passwordResetToken.UserID,
		Password: util.NewNullString(hashedPassword),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update user password: %s", err)
	}

	message := fmt.Sprintf("successfully updated password")

	response := &pb.ResetPasswordResponse{
		Message: message,
		User:    convertUser(user),
	}

	return response, nil
}
