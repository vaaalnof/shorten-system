package analytics

import (
	"context"

	"shortener-service/internal/entity"
	"shortener-service/internal/security"
	"shortener-service/internal/usecase/port"
)

type AnalyticsEventProcessor struct {
	enricher   *AnalyticsEventEnricher
	persister  *AnalyticsEventPersister
	aggregator *AnalyticsEventAggregator
}

func NewAnalyticsEventProcessor(
	analyticsEventRepo port.AnalyticsEventRepository,
	urlDailyAnalyticsRepo port.URLDailyAnalyticsRepository,
	urlDailyVisitorRepo port.URLDailyVisitorRepository,
	visitorHash security.VisitorHash,
	geoIP security.GeoIP,
) *AnalyticsEventProcessor {

	return &AnalyticsEventProcessor{
		enricher: NewAnalyticsEventEnricher(
			geoIP,
		),

		persister: NewAnalyticsEventPersister(
			analyticsEventRepo,
		),

		aggregator: NewAnalyticsEventAggregator(
			urlDailyAnalyticsRepo,
			urlDailyVisitorRepo,
			visitorHash,
		),
	}
}

func (p *AnalyticsEventProcessor) Process(
	ctx context.Context,
	event *entity.AnalyticsEvent,
) error {

	if event == nil {
		return nil
	}

	p.enricher.Enrich(
		event,
	)

	if err := p.persister.Persist(
		ctx,
		event,
	); err != nil {

		return err
	}

	return p.aggregator.Aggregate(
		ctx,
		event,
	)
}
