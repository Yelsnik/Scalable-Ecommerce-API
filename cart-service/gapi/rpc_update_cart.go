package gapi

import (
	pb "cart-service/cart"
	db "cart-service/db/sqlc"
	"cart-service/util"
	"cart-service/val"
	"context"
	"database/sql"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) UpdateCartTx(ctx context.Context, req *pb.UpdateCartTxRequest) (*pb.CartTxResult, error) {
	// validate request
	violations := validateUpdateCartTx(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	// convert request id to uuid
	id, err := util.ConvertStringToUUID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid id %s", err)
	}

	// get the cart item
	cartItem, err := server.store.GetCartitem(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "cart item does not exist %s", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to find cart item %s", err)
	}

	// calculate the subtotal with the new quantity
	subTotal := cartItem.Price * float64(req.Quantity)

	arg := db.UpdateCartitemParams{
		ID:       id,
		Quantity: req.Quantity,
		SubTotal: subTotal,
	}

	// update the cart using db transaction
	result, err := server.store.UpdateCartTx(ctx, cartItem.ID, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update cart item %s", err)
	}

	// return the response
	response := &pb.CartTxResult{
		CartItem: &pb.CartItemResponse{
			Id:        result.CartItem.ID.String(),
			Cart:      result.CartItem.Cart.String(),
			Product:   result.CartItem.Product,
			Quantity:  result.CartItem.Quantity,
			Price:     float32(result.CartItem.Price),
			Currency:  result.CartItem.Currency,
			SubTotal:  float32(result.CartItem.SubTotal),
			CreatedAt: timestamppb.New(result.CartItem.CreatedAt),
		},
		Cart: &pb.Cart{
			Id:         result.Cart.ID.String(),
			UserId:     result.Cart.ID.String(),
			TotalPrice: float32(result.Cart.TotalPrice),
		},
	}

	return response, nil
}

func validateUpdateCartTx(req *pb.UpdateCartTxRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateString(req.GetId(), 1, 100); err != nil {
		violations = append(violations, fielViolation("id", err))
	}

	if err := val.ValidateInt(req.GetQuantity()); err != nil {
		violations = append(violations, fielViolation("quantity", err))
	}

	return violations
}
