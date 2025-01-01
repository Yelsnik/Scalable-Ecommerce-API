package gapi

import (
	"context"
	"fmt"
	db "payment-service/db/sqlc"
	"payment-service/payment"
	"payment-service/util"
	"payment-service/val"

	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/customer"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) StripeCustomer(ctx context.Context, req *payment.StripeCustomerRequest) (*payment.StripeCustomerResponse, error) {

	violations := validateStripeCustomerRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	id, err := util.ConvertStringToUUID(req.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid id %s", err)
	}

	email := fmt.Sprintf("email: '%s'", req.GetEmail())
	searchParams := &stripe.CustomerSearchParams{
		SearchParams: stripe.SearchParams{
			Query: email,
		},
	}

	result := customer.Search(searchParams)
	data := result.CustomerSearchResult().Data

	for i := range data {
		if data[i].Email == req.GetEmail() {
			return nil, nil
		}
	}

	params := &stripe.CustomerParams{
		Email:  stripe.String(req.GetEmail()),
		Source: stripe.String(req.GetToken()),
	}

	customer, err := customer.New(params)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create new stripe customer %s", err)
	}

	_, err = server.store.CreateStripeCustomer(ctx, db.CreateStripeCustomerParams{
		ID:     customer.ID,
		UserID: id,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create customer in the db %s", err)
	}

	response := &payment.StripeCustomerResponse{
		Email: customer.Email,
	}

	return response, nil
}

func validateStripeCustomerRequest(req *payment.StripeCustomerRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateString(req.GetUserId(), 1, 100); err != nil {
		violations = append(violations, fielViolation("user_id", err))
	}

	if err := val.ValidateString(req.GetToken(), 5, 100); err != nil {
		violations = append(violations, fielViolation("token", err))
	}

	if err := val.ValidateEmail(req.GetEmail()); err != nil {
		violations = append(violations, fielViolation("email", err))
	}

	return violations
}
