package gapi

import (
	"context"
	"fmt"
	clientdb "payment-service/client/mock"
	mockdb "payment-service/db/mock"
	db "payment-service/db/sqlc"
	helpermck "payment-service/gapi/helpers/mock"
	"payment-service/payment/payment-service"
	stripemck "payment-service/stripe/mock"
	"payment-service/util"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/stripe/stripe-go/v81"
)

type eqNewCustomerParamsMatcher struct {
	arg *stripe.CustomerParams
}

func (e eqNewCustomerParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(*stripe.CustomerParams)
	if !ok {
		return false
	}

	e.arg.Email = arg.Email
	return reflect.DeepEqual(e.arg.Email, arg.Email)
}

func (e eqNewCustomerParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v", e.arg.Email)
}

func EqNewCustomerParams(arg *stripe.CustomerParams) gomock.Matcher {
	return eqNewCustomerParamsMatcher{arg}
}

type Matcher interface {
	// Matches returns whether x is a match.
	Matches(x interface{}) bool

	// String describes what the matcher matches.
	String() string
}

func randomStripeCustomer() *stripe.Customer {
	return &stripe.Customer{
		ID:    "cus_123",
		Email: util.RandomEmail(),
	}
}

func randomCustomer(customerId string, userId uuid.UUID) db.StripeCustomer {
	return db.StripeCustomer{
		ID:     customerId,
		UserID: userId,
	}
}

func randomPaymentMethod() *stripe.PaymentMethod {
	return &stripe.PaymentMethod{
		ID: "pm_123",
	}
}

func TestStripeCustomer(t *testing.T) {

	stripeCustomer := randomStripeCustomer()
	userId := user()
	email := stripeCustomer.Email
	paymentId := util.RandomString(7)
	paymentMethod := randomPaymentMethod()
	customer := randomCustomer(stripeCustomer.ID, userId)

	testCases := []struct {
		name          string
		req           *payment.StripeCustomerRequest
		buildStubs    func(store *mockdb.MockStore, s *stripemck.MockStripe)
		checkResponse func(t *testing.T, res *payment.StripeCustomerResponse, err error)
	}{
		{
			name: "OK",
			req: &payment.StripeCustomerRequest{
				UserId:    userId.String(),
				PaymentId: paymentId,
				Email:     email,
			},
			buildStubs: func(store *mockdb.MockStore, s *stripemck.MockStripe) {

				params := &stripe.CustomerParams{
					Email: stripe.String(email),
				}
				s.EXPECT().NewCustomer(EqNewCustomerParams(params)).
					Times(1).
					Return(stripeCustomer, nil)

				s.EXPECT().AttachPaymentMethod(gomock.Eq(paymentId), gomock.Any()).
					Times(1).
					Return(paymentMethod, nil)

				arg := db.CreateStripeCustomerParams{
					ID:     stripeCustomer.ID,
					UserID: userId,
				}
				store.EXPECT().CreateStripeCustomer(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(customer, nil)

			},
			checkResponse: func(t *testing.T, res *payment.StripeCustomerResponse, err error) {
				require.NoError(t, err)
				require.Equal(t, res.Email, email)
				require.Equal(t, res.CustomerId, stripeCustomer.ID)
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

			tc.buildStubs(store, str)

			server := newTestServer(t, store, client, helper, str)

			res, err := server.StripeCustomer(context.Background(), tc.req)
			tc.checkResponse(t, res, err)
		})
	}
}
