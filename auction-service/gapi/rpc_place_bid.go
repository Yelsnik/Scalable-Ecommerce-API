package gapi

import (
	"auction-service/auction"
	db "auction-service/db/sqlc"
	"auction-service/util"
	"auction-service/val"
	"auction-service/worker"
	"context"
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/redis/go-redis/v9"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) PlaceBid(ctx context.Context, req *auction.PlaceBidRequest) (*auction.PlaceBidResponse, error) {
	violations := validatePlaceBidReq(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	currentPrice, err := server.getCurrentPrice(ctx, req.GetAuctionId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to validate bidiing amount: %s", err)
	}

	if req.GetAmount() <= float32(currentPrice) {
		return nil, status.Error(codes.InvalidArgument, "amount should be more than previous bid")
	}

	// use transaction to handle placing bids
	userId, err := util.ConvertStringToUUID(req.GetBidderId())
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "You should login: %s", err)
	}

	arg := db.PlaceBidTxParams{
		Ctx:    ctx,
		UserID: userId,
		Amount: float64(req.GetAmount()),
		CurrentPrice: pgtype.Float8{
			Float64: float64(req.GetAmount()),
			Valid:   true,
		},
		AuctionId: req.GetAuctionId(),
		SetRedis: func(auctionId string, currentPrice float64) error {
			err := server.redis.Set(ctx, auctionId, currentPrice, 60000).Err()

			return err
		},
		PublishBid: func(task string, payload any, ctx context.Context) error {
			err = server.rabbitmq.Publish(worker.TaskSendBidUpdates, payload, ctx)

			return err
		},
	}

	_, err = server.store.PlaceBidTx(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to place bid: %s", err)
	}

	response := &auction.PlaceBidResponse{
		Success: true,
		Message: "successfully placed bid",
	}

	return response, nil
}

func validatePlaceBidReq(req *auction.PlaceBidRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateString(req.GetAuctionId(), 2, 100); err != nil {
		violations = append(violations, fielViolation("auction_id", err))
	}

	if err := val.ValidateString(req.GetBidderId(), 2, 100); err != nil {
		violations = append(violations, fielViolation("bidder_id", err))
	}

	return violations
}

func (server *Server) getCurrentPrice(ctx context.Context, auctionId string) (float64, error) {
	priceString, err := server.redis.Get(ctx, auctionId).Result()
	if err != nil && err != redis.Nil {
		return 0, err
	}

	id, err := util.ConvertStringToUUID(auctionId)
	if err != nil {
		return 0, err
	}

	if priceString == "" {
		auction, err := server.store.GetAuction(ctx, id)
		if err != nil {
			return 0, err
		}

		err = server.redis.Set(ctx, auctionId, auction.CurrentPrice, 60000).Err()
		if err != nil {
			return 0, err
		}

		return auction.CurrentPrice, nil
	}

	price, err := strconv.ParseFloat(priceString, 32)
	if err != nil {
		return 0, err
	}

	return price, nil

}
