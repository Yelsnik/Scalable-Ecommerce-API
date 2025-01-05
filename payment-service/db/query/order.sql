-- name: CreateOrders :one
INSERT INTO orders (
  user_name, buyer_id, seller_id, total_price, delivery_address, country, status
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: GetOrdersByBuyerID :one
SELECT * FROM orders
WHERE buyer_id = $1 LIMIT 1;

-- name: GetOrdersBySellerID :one
SELECT * FROM orders
WHERE seller_id = $1 LIMIT 1;

-- name: GetOrdersByID :one
SELECT * FROM orders
WHERE id = $1 LIMIT 1;

-- name: GetOrdersForUpdate :one
SELECT * FROM orders
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: UpdateOrders :one
UPDATE orders
  set status = $2
WHERE id = $1
RETURNING *;