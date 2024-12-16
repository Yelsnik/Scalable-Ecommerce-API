package gapi

import (
	"context"
	"database/sql"
	"fmt"
	db "user-service/db/sqlc"
	"user-service/mail"
	"user-service/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) ForgotPassword(ctx context.Context, req *pb.ForgotPasswordRequest) (*pb.ForgotPasswordResponse, error) {

	user, err := server.store.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "user not found: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to file user: %s", err)
	}

	resetToken, payload, err := server.tokenMaker.CreateToken(user.ID, user.Role, server.config.PasswordResetTokenDuration)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create reset token: %s", err)
	}

	arg := db.CreatePasswordResetTokenParams{
		ID:        payload.ID,
		UserID:    user.ID,
		Token:     resetToken,
		ExpiresAt: payload.ExpiredAt,
	}

	_, err = server.store.CreatePasswordResetToken(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create session: %s", err)
	}

	resetLink := fmt.Sprintf("localhost/v1/forgot-password?token=%s", resetToken)

	sender := mail.NewGmailSender(server.config.EmailSenderName, server.config.EmailSenderAddress, server.config.EmailSenderPassword)

	subject := "Reset your password"

	content := fmt.Sprintf(
		`
	<h1> Reset password link </h1>
	<p> click this link to reset password <a href=%s>link<a/> </p>
	`, resetLink)

	to := []string{user.Email}

	attachFiles := []string{}

	err = sender.SendEmail(subject, content, to, nil, nil, attachFiles)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to send email: %s", err)
	}

	message := fmt.Sprintln("password reset link sent succesfully")

	response := &pb.ForgotPasswordResponse{
		Message: message,
	}

	return response, nil
}
