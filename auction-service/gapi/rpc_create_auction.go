package gapi

import (
	"auction-service/auction"
	db "auction-service/db/sqlc"
	"auction-service/util"
	"auction-service/val"
	"context"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) CreateAuction(ctx context.Context, req *auction.CreateAuctionRequest) (*auction.CreateAuctionResponse, error) {
	violations := validateCreateAuctionReq(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	userId, err := util.ConvertStringToUUID(req.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "You are not logged in: %s", err)
	}

	auctionDB, err := server.store.CreateAuction(ctx, db.CreateAuctionParams{
		ProductID:     req.GetProductId(),
		UserID:        userId,
		StartTime:     req.StartTime.AsTime(),
		EndTime:       req.EndTime.AsTime(),
		StartingPrice: float64(req.StartingPrice),
		CurrentPrice:  float64(req.StartingPrice),
		Status:        "active",
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create auction: %s", err)
	}

	response := &auction.CreateAuctionResponse{
		Auction: &auction.Auction{
			Id:            auctionDB.ID.String(),
			ProductId:     auctionDB.ProductID,
			UserId:        auctionDB.UserID.String(),
			StartTime:     timestamppb.New(auctionDB.StartTime),
			EndTime:       timestamppb.New(auctionDB.EndTime),
			StartingPrice: float32(auctionDB.StartingPrice),
			CurrentPrice:  float32(auctionDB.CurrentPrice),
			Status:        auctionDB.Status,
			WinnerId:      auctionDB.WinnerID.String(),
		},
	}

	return response, nil
}

func validateCreateAuctionReq(req *auction.CreateAuctionRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateString(req.GetProductId(), 2, 100); err != nil {
		violations = append(violations, fielViolation("token", err))
	}

	if err := val.ValidateString(req.GetUserId(), 2, 100); err != nil {
		violations = append(violations, fielViolation("token", err))
	}

	return violations
}
