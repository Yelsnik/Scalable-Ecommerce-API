package gapi

import (
	"context"
	"database/sql"

	db "payment-service/db/sqlc"
	"payment-service/payment/payment-service"
	"payment-service/util"
	"payment-service/val"

	"github.com/google/uuid"
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/customer"
	"github.com/stripe/stripe-go/v81/paymentintent"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreatePayment(ctx context.Context, req *payment.CreatePaymentRequest) (*payment.CreatePaymentResponse, error) {
	// validate request
	violations := validateCreatePaymentRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	// get or create the customer id
	customerId, err := server.getOrCreateCustomer(ctx, req.GetUserId(), req.GetEmail(), req.GetPaymentId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get or create customer %s", err)
	}

	// get the cart item
	cartItem, err := server.client.GetCartItem(ctx, req.GetCartItemId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get cart item %s", err)
	}

	// get the product
	product, err := server.client.GetProductByID(ctx, cartItem.Product)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get product %s", err)
	}

	// get the shop
	shop, err := server.client.GetShopByID(ctx, product.Product.Shop)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get shop %s", err)
	}

	// get the stripe seller account with the seller user id
	sellerAcc, err := server.client.GetStripeSellerAccount(ctx, shop.Shop.ShopOwner)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to make payment at this time %s", err)
	}

	stripeSellerAccountId := sellerAcc.Id

	// calculate platform fee
	platformFee := req.GetAmount() * 10 / 100

	// create payment intent
	params := &stripe.PaymentIntentParams{
		Customer:      stripe.String(customerId),
		PaymentMethod: stripe.String(req.GetPaymentId()),
		Amount:        stripe.Int64(int64(req.GetAmount())),
		Currency:      stripe.String(req.GetCurrency()),

		TransferData: &stripe.PaymentIntentTransferDataParams{
			Destination: stripe.String(stripeSellerAccountId),
		},

		ApplicationFeeAmount: stripe.Int64(int64(platformFee)),

		Metadata: map[string]string{
			"sellerId":        sellerAcc.UserId,
			"cartitemId":      cartItem.Id,
			"buyerId":         req.GetUserId(),
			"deliveryAddress": req.GetDeliveryAddress(),
			"country":         req.GetCountry(),
		},
	}
	params.SetIdempotencyKey(uuid.New().String())

	// check if the user wants to save the card for future use
	if req.SaveCard {
		params.SetupFutureUsage = stripe.String(string(stripe.PaymentIntentSetupFutureUsageOnSession))
	}

	pi, err := paymentintent.New(params)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create payment intent %s", err)
	}

	// save payment in db
	id, err := util.ConvertStringToUUID(req.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid id %s", err)
	}

	_, err = server.store.CreatePayment(ctx, db.CreatePaymentParams{
		ID:       pi.ID,
		Amount:   float64(req.GetAmount()),
		Currency: req.GetCurrency(),
		Status:   string(pi.Status),
		UserID:   id,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create payment %s", err)
	}

	// return the client secret
	return &payment.CreatePaymentResponse{
		ClientSecret: pi.ClientSecret,
		BuyerUserId:  req.GetUserId(),
	}, nil
}

func validateCreatePaymentRequest(req *payment.CreatePaymentRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateString(req.GetUserId(), 1, 100); err != nil {
		violations = append(violations, fielViolation("user_id", err))
	}

	if err := val.ValidateString(req.GetPaymentId(), 5, 100); err != nil {
		violations = append(violations, fielViolation("payment_id", err))
	}

	if err := val.ValidateEmail(req.GetEmail()); err != nil {
		violations = append(violations, fielViolation("email", err))
	}

	if err := val.ValidateFloat(float64(req.GetAmount())); err != nil {
		violations = append(violations, fielViolation("amount", err))
	}

	if err := val.ValidateString(req.GetDeliveryAddress(), 5, 100); err != nil {
		violations = append(violations, fielViolation("delivery_address", err))
	}

	if err := val.ValidateString(req.GetCountry(), 5, 100); err != nil {
		violations = append(violations, fielViolation("country", err))
	}

	if err := val.ValidateString(req.GetCartItemId(), 1, 100); err != nil {
		violations = append(violations, fielViolation("cart_item_id", err))
	}

	if err := val.ValidateString(req.GetCurrency(), 3, 100); err != nil {
		violations = append(violations, fielViolation("currency", err))
	}

	saveCard := req.GetSaveCard()
	if err := val.ValidateBool(&saveCard); err != nil {
		violations = append(violations, fielViolation("save_card", err))
	}

	return violations
}

func (server *Server) getOrCreateCustomer(ctx context.Context, userId, email, paymentId string) (string, error) {

	var err error

	id, err := util.ConvertStringToUUID(userId)
	if err != nil {
		return "", err
	}

	// get the customer id if exists
	customerDB, err := server.store.GetStripeCustomerByUserId(ctx, id)
	if err != nil {
		// create a new customer if not exists
		if err == sql.ErrNoRows {
			params := &stripe.CustomerParams{
				Email:         stripe.String(email),
				PaymentMethod: stripe.String(paymentId),
				Metadata: map[string]string{
					"buyer_id": userId,
				},
			}

			customer, err := customer.New(params)
			if err != nil {
				return "", err
			}

			_, err = server.store.CreateStripeCustomer(ctx, db.CreateStripeCustomerParams{
				ID:     customer.ID,
				UserID: id,
			})
			if err != nil {
				return "", err
			}

			return customer.ID, nil
		}

		return "", err
	}

	return customerDB.ID, nil
}
