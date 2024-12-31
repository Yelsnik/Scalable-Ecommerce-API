package gapi

import (
	"context"
	db "user-service/db/sqlc"
	"user-service/pb"
	"user-service/util"
	"user-service/val"

	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/oauth"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) ExchangeStripeCode(ctx context.Context, req *pb.ExchangeCodeRequest) (*pb.ExchangeCodeResponse, error) {
	violations := validateExchangeStripeCodeReq(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	id, err := util.ConvertStringToUUID(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Invalid id %s", err)
	}

	params := &stripe.OAuthTokenParams{
		GrantType: stripe.String("authorization_code"),
		Code:      stripe.String(req.GetCode()),
	}

	token, err := oauth.New(params)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to exchange code: %s", err)
	}

	stripeAccount, err := server.store.CreateStripeAccount(ctx, db.CreateStripeAccountParams{
		ID:     token.StripeUserID,
		UserID: id,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create stripe account: %s", err)
	}

	return &pb.ExchangeCodeResponse{
		StripeUserId: stripeAccount.ID,
	}, nil
}

func validateExchangeStripeCodeReq(req *pb.ExchangeCodeRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateString(req.GetCode(), 1, 50); err != nil {
		violations = append(violations, fielViolation("token", err))
	}

	return violations
}
