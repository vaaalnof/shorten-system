package nats

import (
	"context"
	"encoding/json"
	"errors"

	gonats "github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"

	"shortener-service/internal/config"
	"shortener-service/internal/entity"
	"shortener-service/internal/usecase/port"

	natsutil "shortener-service/internal/utils/nats"
)

var _ port.AnalyticsConsumer = (*Consumer)(nil)

type Consumer struct {
	js       gonats.JetStreamContext
	settings *config.Settings
	log      *logrus.Logger
}

func NewConsumer(
	js gonats.JetStreamContext,
	settings *config.Settings,
	log *logrus.Logger,
) *Consumer {

	return &Consumer{
		js:       js,
		settings: settings,
		log:      log,
	}
}

func (c *Consumer) Consume(
	ctx context.Context,

	handler func(
		context.Context,
		*entity.AnalyticsEvent,
	) error,
) error {

	sub, err := c.js.PullSubscribe(
		natsutil.AnalyticsSubject,
		natsutil.AnalyticsConsumer,
	)

	if err != nil {

		c.log.WithError(err).
			Error(
				"failed to create consumer subscription",
			)

		return err
	}

	defer func() {

		_ = sub.Unsubscribe()

		c.log.Info(
			"analytics consumer unsubscribed",
		)
	}()

	c.log.WithFields(
		logrus.Fields{
			"subject":  natsutil.AnalyticsSubject,
			"consumer": natsutil.AnalyticsConsumer,
		},
	).Info(
		"analytics consumer started",
	)

	for {

		select {

		case <-ctx.Done():

			c.log.Info(
				"analytics consumer stopped",
			)

			return nil

		default:
		}

		msg, err := sub.Fetch(
			c.settings.NATS.Analytics.FetchBatch,

			gonats.MaxWait(
				c.settings.NATS.Analytics.FetchTimeout,
			),
		)

		if err != nil {

			if errors.Is(
				err,
				context.DeadlineExceeded,
			) {

				continue
			}

			if errors.Is(
				err,
				gonats.ErrTimeout,
			) {

				continue
			}

			c.log.WithError(err).
				Error(
					"failed to fetch analytics messages",
				)

			return err
		}

		if len(msg) > 0 {

			c.log.WithField(
				"batch_size",
				len(msg),
			).Debug(
				"analytics messages fetched",
			)
		}

		for _, msg := range msg {

			var event entity.AnalyticsEvent

			if err := json.Unmarshal(
				msg.Data,
				&event,
			); err != nil {

				c.log.WithError(err).
					Warn(
						"failed to unmarshal analytics message",
					)

				_ = msg.Term()

				continue
			}

			c.log.WithFields(
				logrus.Fields{
					"url_id":     event.UrlID,
					"short_code": event.ShortCode,
				},
			).Debug(
				"analytics message received",
			)

			if err := handler(
				ctx,
				&event,
			); err != nil {

				c.log.WithError(err).
					WithFields(
						logrus.Fields{
							"url_id":     event.UrlID,
							"short_code": event.ShortCode,
						},
					).Warn(
					"analytics processing failed",
				)

				_ = msg.Nak()

				continue
			}

			if err := msg.Ack(); err != nil {

				c.log.WithError(err).
					WithFields(
						logrus.Fields{
							"url_id":     event.UrlID,
							"short_code": event.ShortCode,
						},
					).Error(
					"failed to acknowledge analytics message",
				)

				return err
			}

			c.log.WithFields(
				logrus.Fields{
					"url_id":     event.UrlID,
					"short_code": event.ShortCode,
				},
			).Debug(
				"analytics message acknowledged",
			)
		}
	}
}
