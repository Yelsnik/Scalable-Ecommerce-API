package gapi

import (
	"context"
	"user-service/pb"
	"user-service/val"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) VerifyToken(ctx context.Context, req *pb.VerifyTokenRequest) (*pb.VerifyTokenResponse, error) {
	violations := validateVerifyTokenReq(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	payload, err := server.tokenMaker.VerifyToken(req.GetToken())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not verify token: %s", err)
	}

	response := &pb.VerifyTokenResponse{
		Payload: &pb.Payload{
			Id:        payload.ID.String(),
			UserId:    payload.User_ID.String(),
			Role:      payload.Role,
			IssuedAt:  timestamppb.New(payload.IssuedAt),
			ExpiredAt: timestamppb.New(payload.ExpiredAt),
		},
	}

	return response, nil
}

func validateVerifyTokenReq(req *pb.VerifyTokenRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateStringToken(req.GetToken()); err != nil {
		violations = append(violations, fielViolation("token", err))
	}

	return violations
}
