package db

import (
	"context"
	"time"

	"shortener-service/internal/repository"
	"shortener-service/internal/repository/db/query"
	"shortener-service/internal/usecase/port"
)

var _ port.URLDailyAnalyticsRepository = (*URLDailyAnalyticsRepo)(nil)

type URLDailyAnalyticsRepo struct {
	repo *repository.Repository
}

func NewURLDailyAnalyticsRepo(
	repo *repository.Repository,
) *URLDailyAnalyticsRepo {

	return &URLDailyAnalyticsRepo{
		repo: repo,
	}
}

func (r *URLDailyAnalyticsRepo) AddClick(
	ctx context.Context,
	urlID string,
	analyticsDate int64,
) error {

	now := time.Now().Unix()

	_, err := r.repo.Exec(
		ctx,
		query.URLDailyAnalyticsAddClick,

		urlID,
		analyticsDate,
		now,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *URLDailyAnalyticsRepo) AddUniqueVisitor(
	ctx context.Context,
	urlID string,
	analyticsDate int64,
) error {

	now := time.Now().Unix()

	_, err := r.repo.Exec(
		ctx,
		query.URLDailyAnalyticsAddUniqueVisitor,

		urlID,
		analyticsDate,
		now,
	)

	if err != nil {
		return err
	}

	return nil
}
