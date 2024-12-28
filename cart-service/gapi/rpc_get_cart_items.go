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

// get cart items by cart
func (server *Server) GetCartItemsByCart(ctx context.Context, req *pb.GetCartItemsByCartRequest) (*pb.GetCartItemsByCartResponse, error) {
	violations := validateGetCartItemsByCartRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	// convert request cart id to uuid
	id, err := util.ConvertStringToUUID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid id %s", err)
	}

	cartItems, err := server.store.GetCartitemsByCartID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "cart items do not exist %s", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to find cart items %s", err)
	}

	response := &pb.GetCartItemsByCartResponse{
		CartItem: response(cartItems),
	}

	return response, nil
}

func response(cartItems []db.Cartitem) (responses []*pb.CartItemResponse) {
	for _, v := range cartItems {
		res := &pb.CartItemResponse{
			Id:        v.ID.String(),
			Cart:      v.Cart.String(),
			Product:   v.Product,
			Quantity:  v.Quantity,
			Price:     float32(v.Price),
			Currency:  v.Currency,
			SubTotal:  float32(v.SubTotal),
			CreatedAt: timestamppb.New(v.CreatedAt),
		}
		responses = append(responses, res)
	}

	return responses
}

func validateGetCartItemsByCartRequest(req *pb.GetCartItemsByCartRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateString(req.GetId(), 1, 100); err != nil {
		violations = append(violations, fielViolation("id", err))
	}

	return violations
}

// get cart item by id
func (server *Server) GetCartItem(ctx context.Context, req *pb.GetCartItemByIDRequest) (*pb.CartItemResponse, error) {
	violations := validateGetCartItemsByIDRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	// convert request cart id to uuid
	id, err := util.ConvertStringToUUID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid id %s", err)
	}

	cartItem, err := server.store.GetCartitem(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "cart item do not exist %s", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to find cart item %s", err)
	}

	response := &pb.CartItemResponse{
		Id:        cartItem.ID.String(),
		Cart:      cartItem.Cart.String(),
		Product:   cartItem.Product,
		Quantity:  cartItem.Quantity,
		Price:     float32(cartItem.Price),
		Currency:  cartItem.Currency,
		SubTotal:  float32(cartItem.SubTotal),
		CreatedAt: timestamppb.New(cartItem.CreatedAt),
	}

	return response, nil
}

func validateGetCartItemsByIDRequest(req *pb.GetCartItemByIDRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateString(req.GetId(), 1, 100); err != nil {
		violations = append(violations, fielViolation("id", err))
	}

	return violations
}
