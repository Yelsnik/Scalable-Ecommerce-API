package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"user-service/client"
	"user-service/notification/notification"

	amqp091 "github.com/rabbitmq/amqp091-go"
)

const (
	TaskSendVerifyEmail = "task:sendVerifyEmail"
)

type Amqp interface {
	Publish(task string, v any, ctx context.Context) error
	Consume(ctx context.Context) error
}

type AmqpTask struct {
	ch     *amqp091.Channel
	queue  *amqp091.Queue
	client *client.Client
}

func NewAmqpTask(conn *amqp091.Connection, queuename string, dlx bool, client *client.Client) (Amqp, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	args := amqp091.Table{}
	args["x-message-ttl"] = 10000

	if dlx {
		args["x-dead-letter-exchange"] = queuename + "_dlx"
	}

	q, err := ch.QueueDeclare(
		queuename, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		args,
	)
	if err != nil {
		return nil, err
	}

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		return nil, fmt.Errorf("failed to set Qos: %w", err)
	}

	return &AmqpTask{
		ch:     ch,
		queue:  &q,
		client: client,
	}, nil
}

type Message struct {
	Task    string
	Payload any
}

func (a *AmqpTask) Publish(task string, v any, ctx context.Context) error {
	msg := Message{Task: task, Payload: v}
	bytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	err = a.ch.PublishWithContext(context.Background(), "", a.queue.Name, false, false, amqp091.Publishing{
		DeliveryMode: amqp091.Persistent,
		ContentType:  "text/plain",
		MessageId:    TaskSendVerifyEmail,
		Body:         bytes,
	})
	if err != nil {
		return err
	}

	fmt.Println("Successfully published message")

	return nil
}

func (a *AmqpTask) Consume(ctx context.Context) error {

	go func() {
		_, err := a.client.SendEmail(ctx, &notification.SendEmailRequest{})
		if err != nil {
			log.Fatalf("failed to start consumer via grpc: %s", err)
		}
	}()

	fmt.Println("Successfully requested notification service to start consuming messages")

	return nil

}
