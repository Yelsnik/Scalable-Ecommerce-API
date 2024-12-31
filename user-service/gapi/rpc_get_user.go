package gapi

import (
	"context"
	"database/sql"
	"user-service/pb"
	"user-service/util"
	"user-service/val"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) GetUserByID(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error) {
	violations := validateGetUserByIDReq(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	id, err := util.ConvertStringToUUID(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid id: %s", err)
	}

	user, err := server.store.GetUser(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "user not found: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to find user: %s", err)
	}

	userId := util.ConvertUUIDToString(user.ID)

	response := &pb.GetUserByIdResponse{
		User: &pb.User{
			Id:        userId,
			Name:      user.Name,
			Email:     user.Email,
			Role:      user.Role,
			CreatedAt: timestamppb.New(user.CreatedAt),
		},
	}

	return response, nil
}

func validateGetUserByIDReq(req *pb.GetUserByIdRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateString(req.GetId(), 5, 300); err != nil {
		violations = append(violations, fielViolation("id", err))
	}

	return violations
}
