-- name: CreateStripeCustomer :one
INSERT INTO stripe_customers (
  id, 
  user_id
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetStripeCustomerById :one
SELECT * FROM stripe_customers 
WHERE id = $1 LIMIT 1;