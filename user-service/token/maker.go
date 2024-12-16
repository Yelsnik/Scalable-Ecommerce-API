package token

import (
	"time"

	"github.com/google/uuid"
)

type Maker interface {
	CreateToken(user_id uuid.UUID, role string, duration time.Duration) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}
