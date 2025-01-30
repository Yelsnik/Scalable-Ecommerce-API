package gapi

import (
	"context"
	"payment-service/cart/cart-service"
	clientdb "payment-service/client/mock"
	mockdb "payment-service/db/mock"
	db "payment-service/db/sqlc"
	helpermck "payment-service/gapi/helpers/mock"
	"payment-service/payment/payment-service"
	"payment-service/product/product-service"
	stripemck "payment-service/stripe/mock"
	pb "payment-service/user/user-service"
	"payment-service/util"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/stripe/stripe-go/v81"
)

func user() uuid.UUID {
	return util.Test()
}

func randomRequest(userId uuid.UUID, p db.Payment, ci *cart.CartItemResponse) *payment.CreatePaymentRequest {
	return &payment.CreatePaymentRequest{
		UserId:          userId.String(),
		PaymentId:       p.ID,
		Email:           util.RandomEmail(),
		Amount:          float32(p.Amount),
		DeliveryAddress: util.RandomString(11),
		Country:         util.RandomCountry(),
		CartItemId:      ci.Id,
		Currency:        p.Currency,
		SaveCard:        true,
	}
}

func randomPayment(userId uuid.UUID) db.Payment {
	return db.Payment{
		ID:       "pi_123",
		Amount:   float64(util.RandomMoney()),
		Currency: util.RandomCurrency(),
		Status:   "processing",
		UserID:   userId,
	}
}

func randomCartItem() *cart.CartItemResponse {
	price := float32(util.RandomMoney())
	quantity := util.RandomInt(1, 10)
	subTotal := price * float32(quantity)
	return &cart.CartItemResponse{
		Id:       util.RandomString(7),
		Cart:     util.RandomString(7),
		Product:  util.RandomString(7),
		Quantity: quantity,
		Price:    price,
		Currency: util.RandomCurrency(),
		SubTotal: subTotal,
	}
}

func randomProduct(cart *cart.CartItemResponse, shop *product.ShopResponse) *product.ProductResponse {
	return &product.ProductResponse{
		Product: &product.Product{
			Id:           cart.Product,
			Category:     util.RandomString(8),
			ProductName:  util.RandomString(6),
			Description:  util.RandomString(10),
			Brand:        util.RandomString(6),
			ImageName:    util.RandomString(6),
			CountInStock: util.RandomInt(1, 20),
			Price:        float32(util.RandomMoney()),
			Currency:     util.RandomCurrency(),
			Shop:         shop.Shop.Id,
			Rating:       0,
			IsFeatured:   false,
		},
	}
}

func randomShop() *product.ShopResponse {
	return &product.ShopResponse{
		Shop: &product.Shop{
			Id:          util.RandomString(7),
			Name:        util.RandomString(6),
			Description: util.RandomString(10),
			ImageName:   util.RandomString(6),
			ShopOwner:   util.RandomString(7),
		},
	}
}

func randomStripeSellerAcc(shop *product.ShopResponse) *pb.StripeSellerAccountResponse {
	return &pb.StripeSellerAccountResponse{
		Id:     util.RandomString(7),
		UserId: shop.Shop.ShopOwner,
	}
}

func randomPaymentIntent() *stripe.PaymentIntent {
	return &stripe.PaymentIntent{
		ID:           "pi_123",
		ClientSecret: util.RandomString(11),
		Status:       "processing",
	}
}

// type eqCreatePaymentParamsMatcher struct {
// 	arg db.CreatePaymentParams
// }

// func (e eqCreatePaymentParamsMatcher) Matches(x interface{}) bool {
// 	arg, ok := x.(db.CreatePaymentParams)
// 	if !ok {
// 		return false
// 	}

// 	return reflect.DeepEqual(e.arg, arg)
// }

// func (e eqCreatePaymentParamsMatcher) String() string {
// 	return fmt.Sprintf("matches arg %v", e.arg)
// }

// func EqCreatePaymentParams(arg db.CreatePaymentParams) gomock.Matcher {
// 	return eqCreatePaymentParamsMatcher{arg}
// }

// type Matcher interface {
// 	// Matches returns whether x is a match.
// 	Matches(x interface{}) bool

// 	// String describes what the matcher matches.
// 	String() string
// }

func TestCreatePaymentAPI(t *testing.T) {

	userId := user()
	randomPayment := randomPayment(userId)
	randomCI := randomCartItem()
	randomSh := randomShop()
	randomPro := randomProduct(randomCI, randomSh)
	SellerAcc := randomStripeSellerAcc(randomSh)
	request := randomRequest(userId, randomPayment, randomCI)
	paymentIntent := randomPaymentIntent()

	testCases := []struct {
		name          string
		req           *payment.CreatePaymentRequest
		buildStubs    func(store *mockdb.MockStore, helper *helpermck.MockHelper, client *clientdb.MockClientInterface, s *stripemck.MockStripe)
		checkResponse func(t *testing.T, res *payment.CreatePaymentResponse, err error)
	}{
		{
			name: "OK",
			req: &payment.CreatePaymentRequest{
				UserId:          request.UserId,
				PaymentId:       request.PaymentId,
				Email:           request.Email,
				Amount:          request.Amount,
				DeliveryAddress: request.DeliveryAddress,
				Country:         request.Country,
				CartItemId:      request.CartItemId,
				Currency:        request.Currency,
				SaveCard:        request.SaveCard,
			},
			buildStubs: func(store *mockdb.MockStore, helper *helpermck.MockHelper, client *clientdb.MockClientInterface, s *stripemck.MockStripe) {

				helper.EXPECT().GetOrCreateCustomer(gomock.Any(), gomock.Eq(request.UserId), gomock.Eq(request.Email), gomock.Eq(request.PaymentId)).
					Times(1).
					Return("cus_123", nil)

				client.EXPECT().GetCartItem(gomock.Any(), gomock.Eq(randomCI.Id)).
					Times(1).
					Return(randomCI, nil)

				client.EXPECT().GetProductByID(gomock.Any(), gomock.Eq(randomCI.Product)).
					Times(1).
					Return(randomPro, nil)

				client.EXPECT().GetShopByID(gomock.Any(), gomock.Eq(randomPro.Product.Shop)).
					Times(1).
					Return(randomSh, nil)

				client.EXPECT().GetStripeSellerAccount(gomock.Any(), gomock.Eq(randomSh.Shop.ShopOwner)).
					Times(1).
					Return(SellerAcc, nil)

				s.EXPECT().NewPaymentIntent(gomock.Any()).
					Times(1)

				arg := db.CreatePaymentParams{
					ID:       paymentIntent.ID,
					Amount:   float64(request.Amount),
					Currency: request.Currency,
					Status:   string(paymentIntent.Status),
					UserID:   userId,
				}

				store.EXPECT().
					CreatePayment(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(randomPayment, nil)
			},
			checkResponse: func(t *testing.T, res *payment.CreatePaymentResponse, err error) {
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

			tc.buildStubs(store, helper, client, str)

			server := newTestServer(t, store, client, helper, str)

			res, err := server.CreatePayment(context.Background(), tc.req)
			tc.checkResponse(t, res, err)
		})
	}
}
