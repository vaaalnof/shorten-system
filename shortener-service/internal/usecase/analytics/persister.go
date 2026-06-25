package analytics

import (
	"context"

	"shortener-service/internal/entity"
	"shortener-service/internal/usecase/port"
)

type AnalyticsEventPersister struct {
	repo port.AnalyticsEventRepository
}

func NewAnalyticsEventPersister(
	repo port.AnalyticsEventRepository,
) *AnalyticsEventPersister {

	return &AnalyticsEventPersister{
		repo: repo,
	}
}

func (p *AnalyticsEventPersister) Persist(
	ctx context.Context,
	event *entity.AnalyticsEvent,
) error {

	return p.repo.CreateEvent(
		ctx,
		event,
	)
}
