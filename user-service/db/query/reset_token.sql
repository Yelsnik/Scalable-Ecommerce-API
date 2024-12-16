-- name: CreatePasswordResetToken :one
INSERT INTO password_reset_tokens (
  id, user_id, token, expires_at
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetPasswordResetToken :one
SELECT * FROM password_reset_tokens
WHERE id = $1 LIMIT 1;

-- name: GetPasswordResetTokenByToken :one
SELECT * FROM password_reset_tokens
WHERE token = $1 LIMIT 1;