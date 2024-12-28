package gapi

import (
	pb "cart-service/cart"
	db "cart-service/db/sqlc"
	"cart-service/product/product"
	"cart-service/util"
	"database/sql"
	"fmt"

	"cart-service/val"
	"context"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) AddToCart(ctx context.Context, req *pb.AddtoCartRequest) (*pb.CartTxResult, error) {
	violations := validateAddToCartRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "metadata not found")
	}

	userIds := md.Get("userId")
	if len(userIds) == 0 {
		return nil, status.Error(codes.Unauthenticated, "user_id not found in metadata")
	}

	userId := userIds[0]
	// convert id to uuid
	id, err := util.ConvertStringToUUID(userId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Invalid id %s", err)
	}

	product, err := server.client.GetProductByID(ctx, req.ProductID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get product %s", err)
	}

	// check if cart items exists
	_, err = server.store.GetCartitemByProductID(ctx, req.ProductID)
	if err != nil {
		if err == sql.ErrNoRows {
			// No cart item found; proceed as normal
		} else {
			return nil, status.Error(codes.Internal, "failed to find product")
		}
	} else {
		// If no error, a cart item exists
		return nil, status.Error(codes.InvalidArgument, "Product is already added to cart")
	}

	// get carts or create one for the user if no cart
	cart, err := server.store.GetCartByUserID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			arg := db.CreateCartParams{
				UserID:     id,
				TotalPrice: 0,
			}

			// create cart
			cart, err := server.store.CreateCart(ctx, arg)
			if err != nil {
				return nil, status.Error(codes.Internal, "failed to create cart")
			}

			// create cart item
			result, err := writeResponse(req, cart, product, ctx, server)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "failed to add to cart %s", err)
			}

			return result, nil
		}

		return nil, status.Error(codes.Internal, "failed to find cart")
	}

	// create cart item
	result, err := writeResponse(req, cart, product, ctx, server)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to add to cart %s", err)
	}

	return result, nil

}

func writeResponse(req *pb.AddtoCartRequest, cart db.Cart, product *product.ProductResponse, ctx context.Context, server *Server) (*pb.CartTxResult, error) {
	if req.Quantity == 0 {
		req.Quantity = 1
	}

	if req.Quantity > product.Product.CountInStock {
		message := fmt.Sprintf("The number of %s you want to buy has exceeded the quantity available for sale", product.Product.ProductName)
		return nil, status.Errorf(codes.InvalidArgument, "%s", message)
	}

	subTotal := product.Product.Price * float32(req.Quantity)
	arg := db.CreateCartitemParams{
		Cart:     cart.ID,
		Product:  product.Product.Id.String(),
		Quantity: req.Quantity,
		Price:    float64(product.Product.Price),
		Currency: product.Product.Currency,
		SubTotal: float64(subTotal),
	}

	result, err := server.store.AddToCartTx(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to add to cart %s", err)
	}

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

func validateAddToCartRequest(req *pb.AddtoCartRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateString(req.GetProductID(), 30, 100); err != nil {
		violations = append(violations, fielViolation("product_id", err))
	}

	if err := val.ValidateInt(req.GetQuantity()); err != nil {
		violations = append(violations, fielViolation("quantity", err))
	}

	return violations
}
