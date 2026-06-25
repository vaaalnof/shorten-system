package port

import (
	"context"

	"shortener-service/internal/entity"
)

type AnalyticsProducer interface {
	Publish(
		ctx context.Context,
		event *entity.AnalyticsEvent,
	) error
}
