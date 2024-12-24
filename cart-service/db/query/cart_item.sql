-- name: CreateCartitem :one
INSERT INTO cartitems (
  cart, product, quantity, price, currency, sub_total
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: GetCartitem :one
SELECT * FROM cartitems
WHERE id = $1 LIMIT 1;

-- name: GetCartitemsByCartID :many
SELECT * FROM cartitems
WHERE cart = $1
ORDER BY cart;

-- name: GetCartitemByProductID :one
SELECT * FROM cartitems
WHERE product = $1 LIMIT 1;

-- name: GetCartitemForUpdate :one
SELECT * FROM cartitems
WHERE id = $1
FOR NO KEY UPDATE;

-- name: GetAllCartitems :many
SELECT * FROM cartitems;

-- name: GetALLCartitemsForUpdate :many
SELECT * FROM cartitems
FOR NO KEY UPDATE;

-- name: AddSubtotalPrice :one
SELECT COALESCE(SUM(sub_total), 0)::float AS total
FROM cartitems
WHERE cart = $1;

-- name: ListCartitems :many
SELECT * FROM cartitems
ORDER BY  id
LIMIT $1
OFFSET $2;

-- name: UpdateCartitem :one
UPDATE cartitems
  set quantity = $2,
  sub_total = $3
WHERE id = $1
RETURNING *;

-- name: DeleteCartitem :exec
DELETE FROM cartitems
WHERE id = $1;