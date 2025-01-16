-- name: CreateUser :one
INSERT INTO users (
  name,
  email,
  role,
  is_email_verified,
  password
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users
set name = COALESCE(sqlc.narg(name), name),
  email = COALESCE(sqlc.narg(email), email),
  is_email_verified = COALESCE(sqlc.narg(is_email_verified), is_email_verified),
  password = COALESCE(sqlc.narg(password), password)
WHERE id = sqlc.arg(id)
RETURNING *;