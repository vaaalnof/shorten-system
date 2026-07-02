package mail

import (
	"fmt"
	"net/smtp"

	"auth-service/internal/config"
)

type SMTPMailer struct {
	config *config.SMTPConfig
}

func NewSMTPMailer(
	config *config.SMTPConfig,
) *SMTPMailer {

	return &SMTPMailer{
		config: config,
	}
}

func (m *SMTPMailer) Send(
	to string,
	subject string,
	body string,
) error {

	message := []byte(
		fmt.Sprintf(
			"Subject: %s\r\n"+
				"MIME-version: 1.0;\r\n"+
				"Content-Type: text/html; charset=\"UTF-8\";\r\n\r\n"+
				"%s",
			subject,
			body,
		),
	)

	return smtp.SendMail(
		m.config.Address(),
		m.config.Auth,
		m.config.SenderEmail,
		[]string{to},
		message,
	)
}
