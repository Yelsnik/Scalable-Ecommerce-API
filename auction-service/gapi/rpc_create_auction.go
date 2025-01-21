package gapi

import (
	"auction-service/auction"
	"auction-service/val"
	"context"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func (server *Server) CreateAuction(ctx context.Context, req *auction.CreateAuctionRequest) (*auction.CreateAuctionResponse, error) {
	violations := validateCreateAuctionReq(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	return nil, nil
}

func validateCreateAuctionReq(req *auction.CreateAuctionRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateString(req.GetProductId(), 2, 100); err != nil {
		violations = append(violations, fielViolation("token", err))
	}

	if err := val.ValidateString(req.GetUserId(), 2, 100); err != nil {
		violations = append(violations, fielViolation("token", err))
	}

	return violations
}
