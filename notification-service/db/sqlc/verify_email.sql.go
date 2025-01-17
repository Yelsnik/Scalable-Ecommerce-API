// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: verify_email.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createVerifyEmail = `-- name: CreateVerifyEmail :one
INSERT INTO verify_emails (
  user_id, email, user_name, secret_code
) VALUES (
  $1, $2, $3, $4
) RETURNING id, user_id, email, user_name, secret_code, is_used, expires_at, created_at
`

type CreateVerifyEmailParams struct {
	UserID     uuid.UUID `json:"user_id"`
	Email      string    `json:"email"`
	UserName   string    `json:"user_name"`
	SecretCode string    `json:"secret_code"`
}

func (q *Queries) CreateVerifyEmail(ctx context.Context, arg CreateVerifyEmailParams) (VerifyEmail, error) {
	row := q.db.QueryRowContext(ctx, createVerifyEmail,
		arg.UserID,
		arg.Email,
		arg.UserName,
		arg.SecretCode,
	)
	var i VerifyEmail
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Email,
		&i.UserName,
		&i.SecretCode,
		&i.IsUsed,
		&i.ExpiresAt,
		&i.CreatedAt,
	)
	return i, err
}