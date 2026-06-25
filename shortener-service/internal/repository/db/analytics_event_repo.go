package db

import (
	"context"

	"shortener-service/internal/entity"
	"shortener-service/internal/repository"
	"shortener-service/internal/repository/db/query"
	"shortener-service/internal/usecase/port"
)

var _ port.AnalyticsEventRepository = (*AnalyticsEventRepo)(nil)

type AnalyticsEventRepo struct {
	repo *repository.Repository
}

func NewAnalyticsEventRepo(
	repo *repository.Repository,
) *AnalyticsEventRepo {

	return &AnalyticsEventRepo{
		repo: repo,
	}
}

func (r *AnalyticsEventRepo) CreateEvent(
	ctx context.Context,
	event *entity.AnalyticsEvent,
) error {

	_, err := r.repo.Exec(
		ctx,
		query.AnalyticsEventCreate,

		event.UrlID,
		event.ShortCode,
		event.Referer,
		event.Source,
		event.IPAddress,
		event.UserAgent,
		event.Browser,
		event.OS,
		event.Device,
		event.Country,
		event.ClickedAt,
	)

	return err
}
