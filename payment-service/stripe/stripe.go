package stripe

import (
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/paymentintent"
)

type Stripe interface {
	NewPaymentIntent(params *stripe.PaymentIntentParams) (*stripe.PaymentIntent, error)
}

type StripeClient struct {
}

func NewStripeClient() Stripe {
	return &StripeClient{}
}

func (stripe *StripeClient) NewPaymentIntent(params *stripe.PaymentIntentParams) (*stripe.PaymentIntent, error) {
	return paymentintent.New(params)
}
