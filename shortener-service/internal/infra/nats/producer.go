package nats

import (
	"context"
	"encoding/json"
	"time"

	gonats "github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"

	"shortener-service/internal/entity"
	"shortener-service/internal/usecase/port"
	natsutil "shortener-service/internal/utils/nats"
)

var _ port.AnalyticsProducer = (*Producer)(nil)

type Producer struct {
	js  gonats.JetStreamContext
	log *logrus.Logger
}

func NewProducer(
	js gonats.JetStreamContext,
	log *logrus.Logger,
) *Producer {

	return &Producer{
		js:  js,
		log: log,
	}
}

func (p *Producer) Publish(
	ctx context.Context,
	event *entity.AnalyticsEvent,
) error {

	start := time.Now()

	select {

	case <-ctx.Done():

		p.log.WithError(
			ctx.Err(),
		).Warn(
			"publish cancelled",
		)

		return ctx.Err()

	default:
	}

	payload, err := json.Marshal(
		event,
	)

	if err != nil {

		p.log.WithError(err).
			Error(
				"failed to marshal message",
			)

		return err
	}

	msg := &gonats.Msg{
		Subject: natsutil.AnalyticsSubject,
		Data:    payload,
	}

	ack, err := p.js.PublishMsg(
		msg,
	)

	if err != nil {

		p.log.WithFields(
			logrus.Fields{
				"subject": msg.Subject,
				"size":    len(payload),
			},
		).WithError(
			err,
		).Error(
			"failed to publish message",
		)

		return err
	}

	p.log.WithFields(
		logrus.Fields{
			"subject": msg.Subject,
			"stream":  ack.Stream,
			"seq":     ack.Sequence,
			"size":    len(payload),
			"latency": time.Since(start).Milliseconds(),
		},
	).Debug(
		"message published",
	)

	return nil
}
