-- name: CreateUser :one
INSERT INTO users (
  name,
  email,
  role,
  password
) VALUES (
  $1, $2, $3, $4
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
  password = COALESCE(sqlc.narg(password), password)
WHERE id = sqlc.arg(id)
RETURNING *;