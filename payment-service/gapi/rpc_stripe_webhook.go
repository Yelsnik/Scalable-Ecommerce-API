package gapi

import (
	"context"
	"encoding/json"
	"payment-service/payment/payment-service"
	"payment-service/val"

	"github.com/stripe/stripe-go/v81"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) Webhook(ctx context.Context, req *payment.WebhookRequest) (*payment.WebhookResponse, error) {
	var response *payment.WebhookResponse
	
	violations := validateWebhookRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	webHookSecret := server.config.WebhookSigningKey
	event, err := server.stripe.Webhook(req.GetPayload(), req.GetStripeSignature(),
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
		response, err = server.helpers.HandlePaymentIfSuccesful(ctx, paymentIntent)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to create order %s", err)
		}
	}

	return response, nil
}

// validator func for webhook req
func validateWebhookRequest(req *payment.WebhookRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateString(req.GetPayload(), 1, 100); err != nil {
		violations = append(violations, fielViolation("payload", err))
	}

	if err := val.ValidateString(req.GetStripeSignature(), 1, 100); err != nil {
		violations = append(violations, fielViolation("stripe_signature", err))
	}
	return violations
}
