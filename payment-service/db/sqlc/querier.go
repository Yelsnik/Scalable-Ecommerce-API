// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
)

type Querier interface {
	CreateStripeCustomer(ctx context.Context, arg CreateStripeCustomerParams) (StripeCustomer, error)
	GetStripeCustomerById(ctx context.Context, id string) (StripeCustomer, error)
}

var _ Querier = (*Queries)(nil)
