package port

import (
	"context"

	"shortener-service/internal/entity"
)

type AnalyticsConsumer interface {
	Consume(
		ctx context.Context,
		handler func(
			context.Context,
			*entity.AnalyticsEvent,
		) error,
	) error
}
