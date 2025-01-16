-- name: CreateVerifyEmail :one
INSERT INTO verify_emails (
  user_id, email, user_name, secret_code
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

