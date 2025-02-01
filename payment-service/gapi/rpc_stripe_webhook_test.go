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
)

func TestWebhookApi(t *testing.T) {

	payload := util.RandomString(7)
	stripeSig := util.RandomString(7)

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

			},
			checkResponse: func(t *testing.T, res *payment.WebhookResponse, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, res)
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
