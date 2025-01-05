// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateOrderitems(ctx context.Context, arg CreateOrderitemsParams) (OrderItem, error)
	CreateOrders(ctx context.Context, arg CreateOrdersParams) (Order, error)
	CreatePayment(ctx context.Context, arg CreatePaymentParams) (Payment, error)
	CreateStripeCustomer(ctx context.Context, arg CreateStripeCustomerParams) (StripeCustomer, error)
	GetOrderitemByOrderID(ctx context.Context, orderID uuid.UUID) (OrderItem, error)
	GetOrderitems(ctx context.Context, id uuid.UUID) (OrderItem, error)
	GetOrderitemsByOrderID(ctx context.Context, orderID uuid.UUID) ([]OrderItem, error)
	GetOrderitemsForUpdate(ctx context.Context, id uuid.UUID) (OrderItem, error)
	GetOrdersByBuyerID(ctx context.Context, buyerID uuid.UUID) (Order, error)
	GetOrdersByID(ctx context.Context, id uuid.UUID) (Order, error)
	GetOrdersBySellerID(ctx context.Context, sellerID uuid.UUID) (Order, error)
	GetOrdersForUpdate(ctx context.Context, id uuid.UUID) (Order, error)
	GetPayment(ctx context.Context, id string) (Payment, error)
	GetPaymentByUserID(ctx context.Context, userID uuid.UUID) (Payment, error)
	GetStripeCustomerById(ctx context.Context, id string) (StripeCustomer, error)
	GetStripeCustomerByUserId(ctx context.Context, userID uuid.UUID) (StripeCustomer, error)
	UpdateOrders(ctx context.Context, arg UpdateOrdersParams) (Order, error)
	UpdatePaymentStatus(ctx context.Context, arg UpdatePaymentStatusParams) (Payment, error)
}

var _ Querier = (*Queries)(nil)
