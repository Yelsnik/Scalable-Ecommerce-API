package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	db "notification-service/db/sqlc"
	"notification-service/mail"
	"notification-service/util"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
)

const (
	TaskSendVerifyEmail = "task:sendVerifyEmail"
	TaskSendBidUpdates  = "task:SendBidUpdates"
	QueueBid            = "biddingQueue"
)

type ConsumeTask struct {
	Ctx context.Context
}

type Task interface {
	Consume(queuename string, msgTTL int, dlx bool, arg ConsumeTask) error
}

type PayloadEmail struct {
	Email      string
	UserName   string
	Id         string
	SecretCode string
}

type MessageSendEmail struct {
	Task    string
	Payload PayloadEmail
}

type MessageSendBidUpdate struct {
	Task    string
	Payload BidUpdate
}

type TaskConsumer struct {
	channel *amqp.Channel
	store   db.Store
	mailer  mail.EmailSender
	redis   *redis.Client
}

func NewTaskConsumer(conn *amqp.Connection, store db.Store, mailer mail.EmailSender, redis *redis.Client) (Task, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &TaskConsumer{
		channel: ch,
		store:   store,
		mailer:  mailer,
		redis:   redis,
	}, nil
}

func (tc *TaskConsumer) Consume(queuename string, msgTTL int, dlx bool, arg ConsumeTask) error {
	args := amqp.Table{}

	if msgTTL > 0 {
		args["x-message-ttl"] = msgTTL
	}

	if dlx {
		args["x-dead-letter-exchange"] = queuename + "_dlx"
	}

	_, err := tc.channel.QueueDeclare(
		queuename, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		args,      // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare a queue: %w", err)
	}

	err = tc.channel.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		return fmt.Errorf("failed to set Qos: %w", err)
	}

	messages, err := tc.channel.Consume(
		queuename,
		"",    // consumer
		false, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Graceful shutdown using context
	ctx, cancel := context.WithCancel(arg.Ctx)
	defer cancel()

	go func() {
		for message := range messages {
			if queuename == "email-service" {
				// Process each msg
				var msg MessageSendEmail
				err := json.Unmarshal(message.Body, &msg)
				if err != nil {
					log.Fatalf("fail to unmarshal message: %s", err)
					continue
				}

				// process task send verify email
				if msg.Task == TaskSendVerifyEmail {
					id, err := util.ConvertStringToUUID(msg.Payload.Id)
					if err != nil {
						log.Printf("invalid payload id: %s", err)
						continue
					}

					err = tc.TaskSendVerifyEmail(arg.Ctx, TaskSendVerifyEmailRequest{
						Arg: db.CreateVerifyEmailParams{
							UserID:     id,
							Email:      msg.Payload.Email,
							UserName:   msg.Payload.UserName,
							SecretCode: msg.Payload.SecretCode,
						},
					})
					if err != nil {
						fmt.Printf("failed to handle message: %v", err)
						message.Nack(false, false)
						continue
					}
					message.Ack(false)
				} else {
					fmt.Printf("task handler does not exist for the task in this queue: %v", msg.Task)
					message.Nack(false, false)
				}
			} else if queuename == QueueBid {
				var msg MessageSendBidUpdate
				err = json.Unmarshal(message.Body, &msg)
				if err != nil {
					log.Printf("fail to unmarshal message: %s", err)
					continue
				}

				if msg.Task == TaskSendBidUpdates {
					err = tc.TaskSendBidUpdate(ctx, BidUpdate{
						ID:        msg.Payload.ID,
						UserID:    msg.Payload.UserID,
						AuctionID: msg.Payload.AuctionID,
						Amount:    msg.Payload.Amount,
						BidTime:   msg.Payload.BidTime,
					})

					if err != nil {
						fmt.Printf("failed to handle message: %v", err)
						message.Nack(false, false)
						continue
					}
					message.Ack(false)
				} else {
					fmt.Printf("task handler does not exist for the task in this queue: %v", msg.Task)
					message.Nack(false, false)
				}
			}

		}
	}()

	<-ctx.Done()
	log.Println("shutting down consumer gracefully")

	defer tc.channel.Close()

	return nil
}
