package gapi

import (
	pb "cart-service/cart"
	"cart-service/util"
	"cart-service/val"
	"context"
	"database/sql"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) RemoveCartTx(ctx context.Context, req *pb.RemoveCartTxRequest) (*pb.RemoveCartTxResult, error) {
	// validate request
	violations := validateRemoveCartTx(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	// convert request id to uuid
	id, err := util.ConvertStringToUUID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid id %s", err)
	}

	cartItem, err := server.store.GetCartitem(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "cart item does not exist %s", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to find cart item %s", err)
	}

	result, err := server.store.RemoveCartTx(ctx, cartItem.ID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to remove cart item %s", err)
	}

	response := &pb.RemoveCartTxResult{
		Cart: &pb.Cart{
			Id:         result.Cart.ID.String(),
			UserId:     result.Cart.UserID.String(),
			TotalPrice: float32(result.Cart.TotalPrice),
		},
	}

	return response, nil

}

func validateRemoveCartTx(req *pb.RemoveCartTxRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateString(req.GetId(), 1, 100); err != nil {
		violations = append(violations, fielViolation("id", err))
	}

	return violations
}
