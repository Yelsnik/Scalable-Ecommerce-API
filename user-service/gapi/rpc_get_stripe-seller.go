package gapi

import (
	"context"
	"user-service/pb"
	"user-service/util"
	"user-service/val"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) GetStripeSellerAccount(ctx context.Context, req *pb.StripeSellerAccountRequest) (*pb.StripeSellerAccountResponse, error) {
	violations := validateGetStripeSellerAccReq(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	id, err := util.ConvertStringToUUID(req.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid id %s", err)
	}

	stripeAccount, err := server.store.GetStripeAccountByUserId(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to process payment at this time %s", err)
	}

	response := &pb.StripeSellerAccountResponse{
		Id:        stripeAccount.ID,
		UserId:    stripeAccount.UserID.String(),
		CreatedAt: timestamppb.New(stripeAccount.CreatedAt),
	}

	return response, nil
}

func validateGetStripeSellerAccReq(req *pb.StripeSellerAccountRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateStringToken(req.GetUserId()); err != nil {
		violations = append(violations, fielViolation("token", err))
	}

	return violations
}
