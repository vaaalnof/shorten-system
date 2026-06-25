package port

import "context"

type URLDailyAnalyticsRepository interface {
	AddClick(
		ctx context.Context,
		urlID string,
		analyticsDate int64,
	) error

	AddUniqueVisitor(
		ctx context.Context,
		urlID string,
		analyticsDate int64,
	) error
}
