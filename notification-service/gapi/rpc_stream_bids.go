package gapi

import (
	"encoding/json"
	"log"
	"notification-service/auction"
	"notification-service/worker"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) GetBids(req *auction.GetBidsRequest, stream grpc.ServerStreamingServer[auction.GetBidsResponse]) error {

	ctx := stream.Context()

	// create a redis pubsub channel
	pubsub := server.redis.Subscribe(ctx, req.GetAuctionId())
	defer pubsub.Close()

	// listen for messages on the redis channel
	for {
		select {
		case <-ctx.Done():
			return nil
		case msg := <-pubsub.Channel():
			var bidUpdate worker.BidUpdate
			if err := json.Unmarshal([]byte(msg.Payload), &bidUpdate); err != nil {
				log.Printf("Error unmarshalling update: %v", err)
				continue
			}

			// Send the update to the client
			if err := stream.Send(&auction.GetBidsResponse{
				Bid: &auction.Bid{
					Id:        bidUpdate.ID.String(),
					BidderId:  bidUpdate.UserID.String(),
					AuctionId: bidUpdate.AuctionID.String(),
					Amount:    float32(bidUpdate.Amount),
					BidTime:   timestamppb.New(bidUpdate.BidTime),
				},
			}); err != nil {
				log.Printf("Error sending update to client %s: %v", req.GetUserId(), err)
				return err
			}
		}
	}

}
