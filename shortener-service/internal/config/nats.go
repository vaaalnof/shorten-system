package config

import (
	"time"

	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type NATSConfig struct {
	Conn *nats.Conn
	JS   nats.JetStreamContext
	Log  *logrus.Logger
}

func NewNATSClient(
	v *viper.Viper,
	log *logrus.Logger,
) *NATSConfig {

	url := v.GetString(
		"nats.shorturl",
	)

	timeout := time.Duration(
		v.GetInt(
			"nats.timeout_seconds",
		),
	) * time.Second

	reconnectWait := time.Duration(
		v.GetInt(
			"nats.reconnect_wait_seconds",
		),
	) * time.Second

	maxReconnects := v.GetInt(
		"nats.max_reconnects",
	)

	reconnectBufferMB := v.GetInt(
		"nats.reconnect_buffer_mb",
	)

	conn, err := nats.Connect(
		url,

		nats.Name(
			v.GetString(
				"app.name",
			),
		),

		nats.Timeout(
			timeout,
		),

		nats.MaxReconnects(
			maxReconnects,
		),

		nats.ReconnectWait(
			reconnectWait,
		),

		nats.RetryOnFailedConnect(
			true,
		),

		nats.ReconnectBufSize(
			reconnectBufferMB*1024*1024,
		),

		nats.ErrorHandler(
			func(
				nc *nats.Conn,
				sub *nats.Subscription,
				err error,
			) {

				log.WithError(err).
					Error(
						"nats async error",
					)
			},
		),

		nats.DisconnectErrHandler(
			func(
				nc *nats.Conn,
				err error,
			) {

				log.WithError(err).
					Warn(
						"nats disconnected",
					)
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
			Fatal(
				"failed to connect to nats",
			)
	}

	if err := conn.Flush(); err != nil {

		log.WithError(err).
			Fatal(
				"failed to flush nats connection",
			)
	}

	if err := conn.LastError(); err != nil {

		log.WithError(err).
			Fatal(
				"failed to verify nats connection",
			)
	}

	js, err := conn.JetStream()

	if err != nil {

		log.WithError(err).
			Fatal(
				"failed to create jetstream context",
			)
	}

	log.Infof(
		"connected to nats at %s",
		url,
	)

	return &NATSConfig{
		Conn: conn,
		JS:   js,
		Log:  log,
	}
}
