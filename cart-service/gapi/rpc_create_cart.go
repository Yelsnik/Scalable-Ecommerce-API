package gapi

import (
	pb "cart-service/cart"
	db "cart-service/db/sqlc"
	"cart-service/util"
	"database/sql"
	"fmt"

	"cart-service/val"
	"context"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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

	product, err := server.store.GetProducts(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
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
			writeResponse(req, cart, product, ctx, server)
		}

		return nil, status.Error(codes.Internal, "failed to find cart")
	}

	// create cart item
	writeResponse(req, cart, product, ctx, server)

	return nil, nil
}

func writeResponse(req *pb.AddtoCartRequest, cart db.Cart, product db.Product, ctx context.Context, server *Server) (*pb.CartTxResult, error) {
	if req.Quantity == 0 {
		req.Quantity = 1
	}

	if req.Quantity > product.CountInStock {
		message := fmt.Sprintf("The number of %s you want to buy has exceeded the quantity available for sale", product.ProductName)
		return nil, status.Errorf(codes.InvalidArgument, "%s", message)
	}

	subTotal := product.Price * float64(1)
	arg := db.CreateCartitemParams{
		Cart:     cart.ID,
		Product:  product.ID,
		Quantity: req.Quantity,
		Price:    product.Price,
		Currency: product.Currency,
		SubTotal: subTotal,
	}

	result, err := server.store.AddToCartTx(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to add to cart %s", err)
	}

	response := newCartResponse(result)

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
