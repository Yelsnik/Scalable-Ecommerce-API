package helpers

import (
	"context"
	"database/sql"
	db "payment-service/db/sqlc"
	"payment-service/util"

	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/customer"
)

type Helper interface {
	GetOrCreateCustomer(ctx context.Context, userId string, email string, paymentId string) (string, error)
}

type HelperStruct struct {
	store db.Store
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
