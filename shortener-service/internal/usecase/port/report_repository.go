package port

import (
	"context"

	"shortener-service/internal/entity"
)

type ReportRepository interface {
	GetSummary(
		ctx context.Context,
		urlID string,
	) (
		*entity.ReportSummary,
		error,
	)

	GetChart(
		ctx context.Context,
		urlID string,
	) (
		[]*entity.ReportChart,
		error,
	)

	GetReferrers(
		ctx context.Context,
		urlID string,
	) ([]*entity.ReportReferrer, error)

	GetCountries(
		ctx context.Context,
		urlID string,
	) ([]*entity.ReportCountry, error)

	GetDevices(
		ctx context.Context,
		urlID string,
	) (
		[]*entity.ReportDevice,
		error,
	)
}
