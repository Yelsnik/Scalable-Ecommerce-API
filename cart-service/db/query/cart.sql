-- name: CreateCart :one
INSERT INTO carts (
 user_id, total_price
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetCartByUserID :one
SELECT * FROM carts
WHERE user_id = $1 LIMIT 1;

-- name: GetCart :one
SELECT * FROM carts
WHERE id = $1 LIMIT 1;

-- name: GetCartForUpdate :one
SELECT * FROM carts
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: UpdateCart :one
UPDATE carts
  set total_price = $2
WHERE id = $1
RETURNING *;