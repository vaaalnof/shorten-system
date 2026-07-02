package config

import "github.com/sirupsen/logrus"

func NewLogger(
	cfg LogSettings,
) *logrus.Logger {

	log := logrus.New()

	log.SetLevel(
		cfg.Level,
	)

	log.SetFormatter(
		&logrus.JSONFormatter{},
	)

	return log
}
