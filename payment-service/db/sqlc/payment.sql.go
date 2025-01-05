// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: payment.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createPayment = `-- name: CreatePayment :one
INSERT INTO payments (
  id, amount, currency, status, user_id
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING id, amount, currency, status, user_id, created_at
`

type CreatePaymentParams struct {
	ID       string    `json:"id"`
	Amount   float64   `json:"amount"`
	Currency string    `json:"currency"`
	Status   string    `json:"status"`
	UserID   uuid.UUID `json:"user_id"`
}

func (q *Queries) CreatePayment(ctx context.Context, arg CreatePaymentParams) (Payment, error) {
	row := q.db.QueryRowContext(ctx, createPayment,
		arg.ID,
		arg.Amount,
		arg.Currency,
		arg.Status,
		arg.UserID,
	)
	var i Payment
	err := row.Scan(
		&i.ID,
		&i.Amount,
		&i.Currency,
		&i.Status,
		&i.UserID,
		&i.CreatedAt,
	)
	return i, err
}

const getPayment = `-- name: GetPayment :one
SELECT id, amount, currency, status, user_id, created_at FROM payments
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetPayment(ctx context.Context, id string) (Payment, error) {
	row := q.db.QueryRowContext(ctx, getPayment, id)
	var i Payment
	err := row.Scan(
		&i.ID,
		&i.Amount,
		&i.Currency,
		&i.Status,
		&i.UserID,
		&i.CreatedAt,
	)
	return i, err
}

const getPaymentByUserID = `-- name: GetPaymentByUserID :one
SELECT id, amount, currency, status, user_id, created_at FROM payments
WHERE user_id = $1 LIMIT 1
`

func (q *Queries) GetPaymentByUserID(ctx context.Context, userID uuid.UUID) (Payment, error) {
	row := q.db.QueryRowContext(ctx, getPaymentByUserID, userID)
	var i Payment
	err := row.Scan(
		&i.ID,
		&i.Amount,
		&i.Currency,
		&i.Status,
		&i.UserID,
		&i.CreatedAt,
	)
	return i, err
}

const updatePaymentStatus = `-- name: UpdatePaymentStatus :one
UPDATE payments
  set status = $2
WHERE id = $1
RETURNING id, amount, currency, status, user_id, created_at
`

type UpdatePaymentStatusParams struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

func (q *Queries) UpdatePaymentStatus(ctx context.Context, arg UpdatePaymentStatusParams) (Payment, error) {
	row := q.db.QueryRowContext(ctx, updatePaymentStatus, arg.ID, arg.Status)
	var i Payment
	err := row.Scan(
		&i.ID,
		&i.Amount,
		&i.Currency,
		&i.Status,
		&i.UserID,
		&i.CreatedAt,
	)
	return i, err
}
