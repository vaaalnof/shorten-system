package port

import (
	"context"

	"shortener-service/internal/entity"
)

type AnalyticsEventRepository interface {
	CreateEvent(
		ctx context.Context,
		event *entity.AnalyticsEvent,
	) error
}
