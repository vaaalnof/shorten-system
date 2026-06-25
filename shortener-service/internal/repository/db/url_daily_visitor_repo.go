package db

import (
	"context"
	"time"

	"shortener-service/internal/repository"
	"shortener-service/internal/repository/db/query"
	"shortener-service/internal/usecase/port"
)

var _ port.URLDailyVisitorRepository = (*URLDailyVisitorRepo)(nil)

type URLDailyVisitorRepo struct {
	repo *repository.Repository
}

func NewURLDailyVisitorRepo(
	repo *repository.Repository,
) *URLDailyVisitorRepo {

	return &URLDailyVisitorRepo{
		repo: repo,
	}
}

func (r *URLDailyVisitorRepo) AddVisitor(
	ctx context.Context,
	urlID string,
	analyticsDate int64,
	visitorHash string,
) (bool, error) {

	result, err := r.repo.Exec(
		ctx,
		query.URLDailyVisitorCreate,

		urlID,
		analyticsDate,
		visitorHash,
		time.Now().Unix(),
	)

	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}
