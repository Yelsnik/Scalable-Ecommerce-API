-- name: CreateAuction :one
INSERT INTO auctions (
  product_id,
  user_id,
  start_time,
  end_time,
  starting_price,
  current_price,
  status,
  winner_id
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

-- name: GetAuction :one
SELECT * FROM auctions
WHERE id = $1 LIMIT 1;

-- name: GetAuctionForUpdate :one
SELECT * FROM auctions
WHERE id = $1
FOR NO KEY UPDATE;

-- name: GetAllAuctions :many
SELECT * FROM auctions;

-- name: ListAuctions :many
SELECT * FROM auctions
ORDER BY  id
LIMIT $1
OFFSET $2;

-- name: UpdateAuction :one
UPDATE auctions
set end_time = COALESCE(sqlc.narg(end_time), end_time),
  current_price = COALESCE(sqlc.narg(current_price), current_price),
  status = COALESCE(sqlc.narg(status), status)
WHERE id = sqlc.arg(id)
RETURNING *;