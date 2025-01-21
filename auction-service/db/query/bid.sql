-- name: CreateBid :one
INSERT INTO bids (
  user_id,
  auction_id,
  amount
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetBid :one
SELECT * FROM bids
WHERE id = $1 LIMIT 1;