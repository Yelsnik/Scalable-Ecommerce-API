package stripe

import (
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/paymentintent"
	"github.com/stripe/stripe-go/v81/webhook"
)

type Stripe interface {
	NewPaymentIntent(params *stripe.PaymentIntentParams) (*stripe.PaymentIntent, error)
	Webhook(payload string, header string, secret string) (stripe.Event, error)
}

type StripeClient struct {
}

func NewStripeClient() Stripe {
	return &StripeClient{}
}

func (stripe *StripeClient) NewPaymentIntent(params *stripe.PaymentIntentParams) (*stripe.PaymentIntent, error) {
	return paymentintent.New(params)
}

func (stripe *StripeClient) Webhook(payload string, header string, secret string) (stripe.Event, error) {
	return webhook.ConstructEvent([]byte(payload), header,
		secret)
}
