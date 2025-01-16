package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"user-service/client"
	"user-service/notification/notification"

	amqp091 "github.com/rabbitmq/amqp091-go"
)

const (
	TaskSendVerifyEmail = "task:sendVerifyEmail"
)

type Amqp interface {
	Publish(task string, arg *notification.SendEmailRequest, v any, ctx context.Context) error
}

type AmqpTask struct {
	ch     *amqp091.Channel
	queue  *amqp091.Queue
	client *client.Client
}

func NewAmqpTask(conn *amqp091.Connection, queuename string, client *client.Client) (Amqp, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		queuename, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,
	)
	if err != nil {
		return nil, err
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

func (a *AmqpTask) Publish(task string, arg *notification.SendEmailRequest, v any, ctx context.Context) error {
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

	if msg.Task == TaskSendVerifyEmail {
		_, err = a.client.SendEmail(ctx, arg)
		if err != nil {
			return err
		}

		fmt.Println("Successfully consumed message")
	}

	return nil
}
