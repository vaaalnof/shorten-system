package config

import (
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

type NATSConfig struct {
	Conn *nats.Conn
	JS   nats.JetStreamContext
	Log  *logrus.Logger
}

func NewNATSClient(
	appName string,
	cfg NATSSettings,
	log *logrus.Logger,
) *NATSConfig {

	conn, err := nats.Connect(
		cfg.URL,

		nats.Name(
			appName,
		),

		nats.Timeout(
			cfg.Timeout,
		),

		nats.MaxReconnects(
			cfg.MaxReconnects,
		),

		nats.ReconnectWait(
			cfg.ReconnectWait,
		),

		nats.RetryOnFailedConnect(
			true,
		),

		nats.ReconnectBufSize(
			cfg.ReconnectBufferMB*1024*1024,
		),

		nats.ErrorHandler(
			func(
				nc *nats.Conn,
				sub *nats.Subscription,
				err error,
			) {

				log.WithError(err).
					Error("nats async error")
			},
		),

		nats.DisconnectErrHandler(
			func(
				nc *nats.Conn,
				err error,
			) {

				log.WithError(err).
					Warn("nats disconnected")
			},
		),

		nats.ReconnectHandler(
			func(
				nc *nats.Conn,
			) {

				log.Infof(
					"nats reconnected: %s",
					nc.ConnectedUrl(),
				)
			},
		),

		nats.ClosedHandler(
			func(
				nc *nats.Conn,
			) {

				log.Warn(
					"nats connection closed",
				)
			},
		),
	)

	if err != nil {

		log.WithError(err).
			Fatal("failed to connect to nats")
	}

	if err := conn.Flush(); err != nil {

		log.WithError(err).
			Fatal("failed to flush nats connection")
	}

	if err := conn.LastError(); err != nil {

		log.WithError(err).
			Fatal("failed to verify nats connection")
	}

	js, err := conn.JetStream()

	if err != nil {

		log.WithError(err).
			Fatal("failed to create JetStream context")
	}

	log.Infof(
		"connected to NATS at %s",
		cfg.URL,
	)

	return &NATSConfig{
		Conn: conn,
		JS:   js,
		Log:  log,
	}
}
