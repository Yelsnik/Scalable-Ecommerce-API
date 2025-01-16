package worker

import (
	"context"
	"encoding/json"
	"fmt"
	db "notification-service/db/sqlc"
	"notification-service/mail"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	TaskSendVerifyEmail = "task:sendVerifyEmail"
)

type ConsumeTask struct {
	Ctx  context.Context
	Task TaskSendVerifyEmailRequest
}

type Task interface {
	Consume(queuename string, arg ConsumeTask) error
}

type TaskConsumer struct {
	channel *amqp.Channel
	store   db.Store
	mailer  mail.EmailSender
}

func NewTaskConsumer(conn *amqp.Connection, store db.Store, mailer mail.EmailSender) (Task, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &TaskConsumer{
		channel: ch,
		store:   store,
		mailer:  mailer,
	}, nil
}

func (tc *TaskConsumer) Consume(queuename string, arg ConsumeTask) error {

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

	//forever := make(chan bool)

	go func() {
		for message := range messages {
			// Process each msg
			var msg struct{}

			err := json.Unmarshal(message.Body, &msg)
			if err != nil {
				return
			}

			fmt.Println(msg)
			message.Ack(false)
		}
	}()

	defer tc.channel.Close()

	fmt.Println("Waiting for messages...")

	return nil
}

type TaskSendVerifyEmailRequest struct {
	Arg db.CreateVerifyEmailParams
}

func (tc *TaskConsumer) TaskSendVerifyEmail(ctx context.Context, req TaskSendVerifyEmailRequest) error {
	verifyEmail, err := tc.store.CreateVerifyEmail(ctx, req.Arg)
	if err != nil {
		return fmt.Errorf("failed to create verify email record: %s", err)
	}

	subject := "Welcome to E-commerce app"
	verifyUrl := fmt.Sprintf("localhost:6000?id=%d&secret_code=%s", verifyEmail.ID, verifyEmail.SecretCode)
	content := fmt.Sprintf(`Hello %s <br/>
	Thank you for registering with us! <br/>
	Please <a href="%s">click here</a> to verify your email address, <br/>
	`, verifyEmail.UserName, verifyUrl)
	to := []string{verifyEmail.Email}

	err = tc.mailer.SendEmail(subject, content, to, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to send verify email: %s", err)
	}

	return nil
}
