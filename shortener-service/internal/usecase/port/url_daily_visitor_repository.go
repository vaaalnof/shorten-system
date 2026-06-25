package port

import "context"

type URLDailyVisitorRepository interface {
	AddVisitor(
		ctx context.Context,
		urlID string,
		analyticsDate int64,
		visitorHash string,
	) (bool, error)
}
