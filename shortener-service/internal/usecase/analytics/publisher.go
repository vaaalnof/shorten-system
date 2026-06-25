package analytics

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"

	"shortener-service/internal/delivery/http/middleware"
	"shortener-service/internal/entity"
	"shortener-service/internal/usecase/port"
	"shortener-service/internal/utils/pointer"
)

type AnalyticsEventPublisher struct {
	producer port.AnalyticsProducer
	log      *logrus.Logger
}

func NewAnalyticsEventPublisher(
	producer port.AnalyticsProducer,
	log *logrus.Logger,
) *AnalyticsEventPublisher {

	return &AnalyticsEventPublisher{
		producer: producer,
		log:      log,
	}
}

func (p *AnalyticsEventPublisher) PublishClick(
	meta *middleware.RequestMeta,
	url *entity.URL,
) {

	if p.producer == nil ||
		url == nil {

		return
	}

	event := p.buildEvent(
		meta,
		url,
	)

	publishCtx, cancel := context.WithTimeout(
		context.Background(),
		3*time.Second,
	)

	defer cancel()

	if err := p.producer.Publish(
		publishCtx,
		event,
	); err != nil {

		p.log.WithError(
			err,
		).WithField(
			"url_id",
			url.ID,
		).WithField(
			"short_code",
			url.ShortCode,
		).Warn(
			"failed to publish analytics event",
		)

		return
	}

	p.log.WithField(
		"url_id",
		url.ID,
	).WithField(
		"short_code",
		url.ShortCode,
	).WithField(
		"source",
		pointer.Value(
			event.Source,
		),
	).WithField(
		"browser",
		pointer.Value(
			event.Browser,
		),
	).WithField(
		"os",
		pointer.Value(
			event.OS,
		),
	).WithField(
		"device",
		pointer.Value(
			event.Device,
		),
	).WithField(
		"ip",
		pointer.Value(
			event.IPAddress,
		),
	).Debug(
		"analytics event published",
	)
}
