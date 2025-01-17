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
)

const (
	TaskSendVerifyEmail = "task:sendVerifyEmail"
)

type ConsumeTask struct {
	Ctx context.Context
}

type Task interface {
	Consume(queuename string, msgTTL int, dlx bool, arg ConsumeTask) error
}

type Payload struct {
	Email      string
	UserName   string
	Id         string
	SecretCode string
}

type Message struct {
	Task    string
	Payload Payload
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
			// Process each msg
			var msg Message
			err := json.Unmarshal(message.Body, &msg)
			if err != nil {
				log.Fatalf("fail to unmarshal message: %s", err)
				continue
			}

			// process task send verify email
			if msg.Task == TaskSendVerifyEmail {
				id, err := util.ConvertStringToUUID(msg.Payload.Id)
				if err != nil {
					log.Fatalf("invalid payload id: %s", err)
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
				fmt.Printf("task does not exist: %v", msg.Task)
				message.Nack(false, false)
			}

		}
	}()

	<-ctx.Done()
	log.Println("shutting down consumer gracefully")

	defer tc.channel.Close()

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

	//fmt.Println(subject, content, to)

	err = tc.mailer.SendEmail(subject, content, to, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to send verify email: %s", err)
	}

	return nil
}
