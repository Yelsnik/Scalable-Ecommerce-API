package helpers

import (
	"context"
	"database/sql"
	"payment-service/cart/cart-service"
	"payment-service/client"
	db "payment-service/db/sqlc"
	"payment-service/payment/payment-service"
	"payment-service/product/product-service"
	"payment-service/util"

	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/customer"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Helper interface {
	GetOrCreateCustomer(ctx context.Context, userId string, email string, paymentId string) (string, error)
	HandlePaymentIfSuccesful(ctx context.Context, paymentIntent *stripe.PaymentIntent) (*payment.WebhookResponse, error)
}

type HelperStruct struct {
	store  db.Store
	client client.ClientInterface
}

func NewHelpers() Helper {
	return &HelperStruct{}
}

func (helper *HelperStruct) GetOrCreateCustomer(ctx context.Context, userId, email, paymentId string) (string, error) {

	var err error

	id, err := util.ConvertStringToUUID(userId)
	if err != nil {
		return "", err
	}

	// get the customer id if exists
	customerDB, err := helper.store.GetStripeCustomerByUserId(ctx, id)
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

			_, err = helper.store.CreateStripeCustomer(ctx, db.CreateStripeCustomerParams{
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

func (helper *HelperStruct) HandlePaymentIfSuccesful(ctx context.Context, paymentIntent *stripe.PaymentIntent) (*payment.WebhookResponse, error) {
	// get the customer
	params := &stripe.CustomerParams{}

	customer, err := customer.Get(paymentIntent.Customer.ID, params)
	if err != nil {
		return nil, err
	}

	// get user from db
	user, err := helper.client.GetUserByEmail(ctx, customer.Email)
	if err != nil {
		return nil, err
	}

	// create the order
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
		GetCartItem: func(ctx context.Context, cartItemId string) (*cart.CartItemResponse, error) {
			response, err := helper.client.GetCartItem(ctx, cartItemId)

			return response, err
		},
		GetProductByID: func(ctx context.Context, productId string) (*product.ProductResponse, error) {
			response, err := helper.client.GetProductByID(ctx, productId)

			return response, err
		},
		RemoveCartTx: func(ctx context.Context, cartItemId string) (*cart.RemoveCartTxResult, error) {
			response, err := helper.client.RemoveCartTx(ctx, cartItemId)

			return response, err
		},
	}

	result, err := helper.store.CreateOrderTx(ctx, arg)
	if err != nil {
		return nil, err
	}

	// send the response
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
