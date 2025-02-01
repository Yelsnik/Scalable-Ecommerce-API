package db

import (
	"auction-service/util"
	"auction-service/worker"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

// update auction price in db
// create a new bid in the db
// update redis cache

type PlaceBidTxResult struct {
	Bid     Bid     `json:"bid"`
	Auction Auction `json:"auction"`
}

type PlaceBidTxParams struct {
	Ctx          context.Context
	UserID       uuid.UUID
	Amount       float64
	CurrentPrice pgtype.Float8
	AuctionId    string
	SetRedis     func(auctionId string, currentPrice float64) error
	PublishBid   func(task string, payload any, ctx context.Context) error
}

func (store *SQLStore) PlaceBidTx(ctx context.Context, arg PlaceBidTxParams) (PlaceBidTxResult, error) {
	var result PlaceBidTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		id, err := util.ConvertStringToUUID(arg.AuctionId)
		if err != nil {
			return err
		}

		// get auction for update
		result.Auction, err = q.GetAuctionForUpdate(ctx, id)
		if err != nil {
			return err
		}

		// update auction
		result.Auction, err = q.UpdateAuction(ctx, UpdateAuctionParams{
			CurrentPrice: arg.CurrentPrice,
			ID:           id,
		})
		if err != nil {
			return err
		}

		// insert bid
		result.Bid, err = q.CreateBid(ctx, CreateBidParams{
			UserID:    arg.UserID,
			AuctionID: id,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		// update redis cache
		err = arg.SetRedis(arg.AuctionId, result.Auction.CurrentPrice)
		if err != nil {
			return err
		}

		// publish bid
		err = arg.PublishBid(worker.TaskSendBidUpdates, result.Bid, arg.Ctx)

		return err
	})

	return result, err
}
