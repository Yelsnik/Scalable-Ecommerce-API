package gapi

import (
	"context"
	clientdb "payment-service/client/mock"
	mockdb "payment-service/db/mock"
	helpermck "payment-service/gapi/helpers/mock"
	"payment-service/payment/payment-service"
	stripemck "payment-service/stripe/mock"
	"payment-service/util"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stripe/stripe-go/v81"
)

func randomEvent() stripe.Event {
	return stripe.Event{
		Type: "payment_intent.succeeded",
	}
}

func randomWebhookResponse() *payment.WebhookResponse {
	shop := randomShop()
	sellerAcc := randomStripeSellerAcc(shop)
	user := user()
	p := randomPayment(user)
	cartItem := randomCartItem()
	req := randomRequest(user, p, cartItem)
	return &payment.WebhookResponse{
		Payment: &payment.Payment{
			Id:       p.ID,
			Amount:   float32(p.Amount),
			Currency: p.Currency,
			Status:   p.Status,
			UserId:   p.UserID.String(),
		},
		Order: &payment.Order{
			Id:              util.RandomString(7),
			UserName:        util.RandomString(6),
			BuyerId:         p.UserID.String(),
			SellerId:        sellerAcc.UserId,
			TotalPrice:      float32(p.Amount),
			DeliveryAddress: req.DeliveryAddress,
			Country:         req.Country,
			Status:          "succeeded",
		},
	}
}

func TestWebhookApi(t *testing.T) {

	payload := util.RandomString(7)
	stripeSig := util.RandomString(7)
	event := randomEvent()
	webhookResponse := randomWebhookResponse()

	testCases := []struct {
		name          string
		req           *payment.WebhookRequest
		buildStubs    func(store *mockdb.MockStore, helper *helpermck.MockHelper, s *stripemck.MockStripe)
		checkResponse func(t *testing.T, res *payment.WebhookResponse, err error)
	}{
		{
			name: "OK",
			req: &payment.WebhookRequest{
				Payload:         payload,
				StripeSignature: stripeSig,
			},
			buildStubs: func(store *mockdb.MockStore, helper *helpermck.MockHelper, s *stripemck.MockStripe) {
				s.EXPECT().Webhook(payload, stripeSig, gomock.Any()).
					Times(1).
					Return(event, nil)

				//paymentIntent := &stripe.PaymentIntent{}
				helper.EXPECT().HandlePaymentIfSuccesful(gomock.Any(), gomock.Any()).
					Times(1).
					Return(webhookResponse, nil)
			},
			checkResponse: func(t *testing.T, res *payment.WebhookResponse, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, res)
				require.NotEmpty(t, res.Payment)
				require.NotEmpty(t, res.Order)
			},
		},
		{
			name: "InvalidArgument",
			req: &payment.WebhookRequest{
				Payload:         "",
				StripeSignature: "",
			},
			buildStubs: func(store *mockdb.MockStore, helper *helpermck.MockHelper, s *stripemck.MockStripe) {
				s.EXPECT().Webhook(payload, stripeSig, gomock.Any()).
					Times(0).
					Return(event, nil)

				//paymentIntent := &stripe.PaymentIntent{}
				helper.EXPECT().HandlePaymentIfSuccesful(gomock.Any(), gomock.Any()).
					Times(0).
					Return(webhookResponse, nil)
			},
			checkResponse: func(t *testing.T, res *payment.WebhookResponse, err error) {
				//require.Error(t, status.Code(), )
				require.Nil(t, res)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)

			ctrlClient := gomock.NewController(t)
			defer ctrl.Finish()

			client := clientdb.NewMockClientInterface(ctrlClient)

			ctrlHelper := gomock.NewController(t)
			defer ctrl.Finish()

			helper := helpermck.NewMockHelper(ctrlHelper)

			ctrlStripe := gomock.NewController(t)
			defer ctrl.Finish()

			str := stripemck.NewMockStripe(ctrlStripe)

			tc.buildStubs(store, helper, str)

			server := newTestServer(t, store, client, helper, str)

			res, err := server.Webhook(context.Background(), tc.req)
			tc.checkResponse(t, res, err)
		})
	}
}
