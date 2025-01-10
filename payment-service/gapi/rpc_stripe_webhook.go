package gapi

import (
	"context"
	"encoding/json"
	db "payment-service/db/sqlc"
	"payment-service/payment/payment-service"
	"payment-service/val"

	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/customer"
	"github.com/stripe/stripe-go/v81/webhook"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) Webhook(ctx context.Context, req *payment.WebhookRequest) (*payment.WebhookResponse, error) {
	violations := validateWebhookRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	webHookSecret := server.config.WebhookSigningKey
	event, err := webhook.ConstructEvent([]byte(req.GetPayload()), req.GetStripeSignature(),
		webHookSecret)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to construct stripe event %s", err)
	}

	if event.Type == "payment_intent.payment_failed" {
		var paymentIntent *stripe.PaymentIntent
		err := json.Unmarshal(event.Data.Raw, &paymentIntent)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to unmarshal payment intent %s", err)
		}

		return nil, status.Error(codes.Internal, "payment failed")
	}

	if event.Type == "payment_intent.succeeded" {
		var paymentIntent *stripe.PaymentIntent
		response, err := server.handlePaymentIfSuccesful(ctx, paymentIntent)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to create order %s", err)
		}

		return response, nil
	}

	return &payment.WebhookResponse{}, nil
}

func (server *Server) handlePaymentIfSuccesful(ctx context.Context, paymentIntent *stripe.PaymentIntent) (*payment.WebhookResponse, error) {
	params := &stripe.CustomerParams{}

	customer, err := customer.Get(paymentIntent.Customer.ID, params)
	if err != nil {
		return nil, err
	}

	user, err := server.client.GetUserByEmail(ctx, customer.Email)
	if err != nil {
		return nil, err
	}

	arg := db.OrderTxParams{
		PaymentIntent:   paymentIntent.ID,
		UserName:        user.User.Name,
		BuyerID:         paymentIntent.Metadata["buyerId"],
		SellerID:        paymentIntent.Metadata["sellerId"],
		CartItemId:      paymentIntent.Metadata["cartitemId"],
		TotalPrice:      float64(paymentIntent.AmountReceived),
		DeliveryAddress: paymentIntent.Metadata["deliveryAddress"],
		Country:         paymentIntent.Metadata["country"],
		PaymentStatus:   string(paymentIntent.Status),
		OrderStatus:     "processing",
	}

	result, err := server.store.CreateOrderTx(ctx, arg)
	if err != nil {
		return nil, err
	}

	response := &payment.WebhookResponse{
		Payment: &payment.Payment{
			Id:        result.Payment.ID,
			Amount:    float32(result.Payment.Amount),
			Currency:  result.Payment.Currency,
			Status:    result.Payment.Status,
			UserId:    result.Payment.UserID.String(),
			CreatedAt: timestamppb.New(result.Payment.CreatedAt),
		},
		Order: &payment.Order{
			Id:              result.Order.ID.String(),
			UserName:        result.Order.UserName,
			BuyerId:         result.Order.BuyerID.String(),
			SellerId:        result.Order.SellerID.String(),
			TotalPrice:      float32(result.Order.TotalPrice),
			DeliveryAddress: result.Order.DeliveryAddress,
			Country:         result.Order.Country,
			Status:          result.Order.Status,
			CreatedAt:       timestamppb.New(result.Order.CreatedAt),
		},
	}

	return response, nil

}

func validateWebhookRequest(req *payment.WebhookRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateString(req.GetPayload(), 1, 100); err != nil {
		violations = append(violations, fielViolation("payload", err))
	}

	if err := val.ValidateString(req.GetStripeSignature(), 1, 100); err != nil {
		violations = append(violations, fielViolation("stripe_signature", err))
	}
	return violations
}
