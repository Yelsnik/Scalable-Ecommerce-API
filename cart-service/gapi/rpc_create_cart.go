package gapi

import (
	pb "cart-service/cart"
	//db "cart-service/db/sqlc"
	"cart-service/val"
	"context"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func (server *Server) AddToCart(ctx context.Context, req *pb.AddtoCartRequest) (*pb.CartTxResult, error) {
	violations := validateAddToCartRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	return nil, nil
}

func validateAddToCartRequest(req *pb.AddtoCartRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateString(req.GetProductID(), 30, 100); err != nil {
		violations = append(violations, fielViolation("product_id", err))
	}

	if err := val.ValidateInt(req.GetQuantity()); err != nil {
		violations = append(violations, fielViolation("quantity", err))
	}

	return violations
}
