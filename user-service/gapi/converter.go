package gapi

import (
	db "user-service/db/sqlc"
	"user-service/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.User) *pb.User {
	return &pb.User{
		Id:              user.ID.String(),
		Name:            user.Name,
		Email:           user.Email,
		Role:            user.Role,
		IsEmailVerified: user.IsEmailVerified,
		CreatedAt:       timestamppb.New(user.CreatedAt),
	}
}
