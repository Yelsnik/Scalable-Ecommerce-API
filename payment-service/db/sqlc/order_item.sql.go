// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: order_item.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createOrderitems = `-- name: CreateOrderitems :one
INSERT INTO order_items (
  item_name, item_sub_total, quantity, item_id, order_id
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING id, item_name, item_sub_total, quantity, item_id, order_id, created_at
`

type CreateOrderitemsParams struct {
	ItemName     string    `json:"item_name"`
	ItemSubTotal float64   `json:"item_sub_total"`
	Quantity     int64     `json:"quantity"`
	ItemID       uuid.UUID `json:"item_id"`
	OrderID      uuid.UUID `json:"order_id"`
}

func (q *Queries) CreateOrderitems(ctx context.Context, arg CreateOrderitemsParams) (OrderItem, error) {
	row := q.db.QueryRowContext(ctx, createOrderitems,
		arg.ItemName,
		arg.ItemSubTotal,
		arg.Quantity,
		arg.ItemID,
		arg.OrderID,
	)
	var i OrderItem
	err := row.Scan(
		&i.ID,
		&i.ItemName,
		&i.ItemSubTotal,
		&i.Quantity,
		&i.ItemID,
		&i.OrderID,
		&i.CreatedAt,
	)
	return i, err
}

const getOrderitemByOrderID = `-- name: GetOrderitemByOrderID :one
SELECT id, item_name, item_sub_total, quantity, item_id, order_id, created_at FROM order_items
WHERE order_id = $1 LIMIT 1
`

func (q *Queries) GetOrderitemByOrderID(ctx context.Context, orderID uuid.UUID) (OrderItem, error) {
	row := q.db.QueryRowContext(ctx, getOrderitemByOrderID, orderID)
	var i OrderItem
	err := row.Scan(
		&i.ID,
		&i.ItemName,
		&i.ItemSubTotal,
		&i.Quantity,
		&i.ItemID,
		&i.OrderID,
		&i.CreatedAt,
	)
	return i, err
}

const getOrderitems = `-- name: GetOrderitems :one
SELECT id, item_name, item_sub_total, quantity, item_id, order_id, created_at FROM order_items
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetOrderitems(ctx context.Context, id uuid.UUID) (OrderItem, error) {
	row := q.db.QueryRowContext(ctx, getOrderitems, id)
	var i OrderItem
	err := row.Scan(
		&i.ID,
		&i.ItemName,
		&i.ItemSubTotal,
		&i.Quantity,
		&i.ItemID,
		&i.OrderID,
		&i.CreatedAt,
	)
	return i, err
}

const getOrderitemsByOrderID = `-- name: GetOrderitemsByOrderID :many
SELECT id, item_name, item_sub_total, quantity, item_id, order_id, created_at FROM order_items
WHERE order_id = $1 
ORDER BY order_id
`

func (q *Queries) GetOrderitemsByOrderID(ctx context.Context, orderID uuid.UUID) ([]OrderItem, error) {
	rows, err := q.db.QueryContext(ctx, getOrderitemsByOrderID, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []OrderItem
	for rows.Next() {
		var i OrderItem
		if err := rows.Scan(
			&i.ID,
			&i.ItemName,
			&i.ItemSubTotal,
			&i.Quantity,
			&i.ItemID,
			&i.OrderID,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getOrderitemsForUpdate = `-- name: GetOrderitemsForUpdate :one
SELECT id, item_name, item_sub_total, quantity, item_id, order_id, created_at FROM order_items
WHERE id = $1
FOR NO KEY UPDATE
`

func (q *Queries) GetOrderitemsForUpdate(ctx context.Context, id uuid.UUID) (OrderItem, error) {
	row := q.db.QueryRowContext(ctx, getOrderitemsForUpdate, id)
	var i OrderItem
	err := row.Scan(
		&i.ID,
		&i.ItemName,
		&i.ItemSubTotal,
		&i.Quantity,
		&i.ItemID,
		&i.OrderID,
		&i.CreatedAt,
	)
	return i, err
}
