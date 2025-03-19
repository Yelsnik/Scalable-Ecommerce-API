package stripe

import (
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/customer"
	"github.com/stripe/stripe-go/v81/paymentintent"
	"github.com/stripe/stripe-go/v81/paymentmethod"
	"github.com/stripe/stripe-go/v81/webhook"
)

type Stripe interface {
	NewPaymentIntent(params *stripe.PaymentIntentParams) (*stripe.PaymentIntent, error)
	Webhook(payload string, header string, secret string) (stripe.Event, error)
	NewCustomer(params *stripe.CustomerParams) (*stripe.Customer, error)
	AttachPaymentMethod(id string, params *stripe.PaymentMethodAttachParams) (*stripe.PaymentMethod, error)
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

func (stripe *StripeClient) NewCustomer(params *stripe.CustomerParams) (*stripe.Customer, error) {
	return customer.New(params)
}

func (stripe *StripeClient) AttachPaymentMethod(id string, params *stripe.PaymentMethodAttachParams) (*stripe.PaymentMethod, error) {
	return paymentmethod.Attach(id, params)
}
