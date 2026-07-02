package config

import (
	"fmt"
	"net/smtp"

	"github.com/sirupsen/logrus"
)

type SMTPConfig struct {
	Host        string
	Port        int
	Username    string
	Password    string
	SenderName  string
	SenderEmail string

	Auth smtp.Auth
	Log  *logrus.Logger
}

func NewSMTPConfig(
	cfg SMTPSettings,
	log *logrus.Logger,
) *SMTPConfig {

	auth := smtp.PlainAuth(
		"",
		cfg.Username,
		cfg.Password,
		cfg.Host,
	)

	log.Infof(
		"SMTP configured (%s:%d)",
		cfg.Host,
		cfg.Port,
	)

	return &SMTPConfig{
		Host:        cfg.Host,
		Port:        cfg.Port,
		Username:    cfg.Username,
		Password:    cfg.Password,
		SenderName:  cfg.SenderName,
		SenderEmail: cfg.SenderEmail,
		Auth:        auth,
		Log:         log,
	}
}

func (c *SMTPConfig) Address() string {

	return fmt.Sprintf(
		"%s:%d",
		c.Host,
		c.Port,
	)
}
