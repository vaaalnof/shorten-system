package analytics

import (
	"context"
	"time"

	"shortener-service/internal/entity"
	"shortener-service/internal/security"
	"shortener-service/internal/usecase/port"
	"shortener-service/internal/utils/pointer"
)

type AnalyticsEventAggregator struct {
	urlDailyAnalyticsRepo port.URLDailyAnalyticsRepository
	urlDailyVisitorRepo   port.URLDailyVisitorRepository
	visitorHash           security.VisitorHash
}

func NewAnalyticsEventAggregator(
	urlDailyAnalyticsRepo port.URLDailyAnalyticsRepository,
	urlDailyVisitorRepo port.URLDailyVisitorRepository,
	visitorHash security.VisitorHash,
) *AnalyticsEventAggregator {

	return &AnalyticsEventAggregator{
		urlDailyAnalyticsRepo: urlDailyAnalyticsRepo,
		urlDailyVisitorRepo:   urlDailyVisitorRepo,
		visitorHash:           visitorHash,
	}
}

func (a *AnalyticsEventAggregator) Aggregate(
	ctx context.Context,
	event *entity.AnalyticsEvent,
) error {

	analyticsDate := time.
		Unix(
			event.ClickedAt,
			0,
		).
		UTC().
		Truncate(
			24 * time.Hour,
		).
		Unix()

	if err := a.urlDailyAnalyticsRepo.AddClick(
		ctx,
		event.UrlID,
		analyticsDate,
	); err != nil {

		return err
	}

	if event.IPAddress == nil {
		return nil
	}

	visitorHash := a.visitorHash.Hash(
		*event.IPAddress,
		pointer.Value(
			event.UserAgent,
		),
	)

	isNewVisitor, err := a.urlDailyVisitorRepo.AddVisitor(
		ctx,
		event.UrlID,
		analyticsDate,
		visitorHash,
	)

	if err != nil {
		return err
	}

	if !isNewVisitor {
		return nil
	}

	return a.urlDailyAnalyticsRepo.AddUniqueVisitor(
		ctx,
		event.UrlID,
		analyticsDate,
	)
}
