-- name: CreateStripeAccount :one
INSERT INTO stripe_accounts (
  id, 
  user_id
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetStripeAccountById :one
SELECT * FROM stripe_accounts 
WHERE id = $1 LIMIT 1;

-- name: GetStripeAccountByUserId :one
SELECT * FROM stripe_accounts 
WHERE user_id = $1 LIMIT 1;