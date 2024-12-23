// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	ID         uuid.UUID `json:"id"`
	UserID     uuid.UUID `json:"user_id"`
	TotalPrice float64   `json:"total_price"`
}

type Cartitem struct {
	ID       uuid.UUID `json:"id"`
	Cart     uuid.UUID `json:"cart"`
	Product  string    `json:"product"`
	Quantity int64     `json:"quantity"`
	// must be positive
	Price     float64   `json:"price"`
	Currency  string    `json:"currency"`
	SubTotal  float64   `json:"sub_total"`
	CreatedAt time.Time `json:"created_at"`
}
