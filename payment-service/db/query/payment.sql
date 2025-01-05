-- name: CreatePayment :one
INSERT INTO payments (
  id, amount, currency, status, user_id
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetPaymentByUserID :one
SELECT * FROM payments
WHERE user_id = $1 LIMIT 1;

-- name: GetPayment :one
SELECT * FROM payments
WHERE id = $1 LIMIT 1;

-- name: UpdatePaymentStatus :one
UPDATE payments
  set status = $2
WHERE id = $1
RETURNING *;