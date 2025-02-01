package worker

import (
	"context"
	"fmt"
	db "notification-service/db/sqlc"
)

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
