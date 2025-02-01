package worker

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type BidUpdate struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	AuctionID uuid.UUID
	Amount    float64
	BidTime   time.Time
}

func (tc *TaskConsumer) TaskSendBidUpdate(ctx context.Context, arg BidUpdate) error {

	bid, err := json.Marshal(arg)
	if err != nil {
		return err
	}

	err = tc.redis.Publish(ctx, arg.AuctionID.String(), bid).Err()

	return err
}
