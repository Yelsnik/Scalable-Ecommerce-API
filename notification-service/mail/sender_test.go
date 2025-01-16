package mail

import (
	"notification-service/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSendEmail(t *testing.T) {

	config, err := util.LoadConfig("..")

	require.NoError(t, err)

	sender := NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)

	subject := "A test email"
	content := `
		<h1>Hello World </h1>
		<p> This is a test meassage</p>
		`

	to := []string{"kingsleyokgeorge@gmail.com"}
	attachFiles := []string{}

	err = sender.SendEmail(subject, content, to, nil, nil, attachFiles)
	//require.NoError(t, err)

}
