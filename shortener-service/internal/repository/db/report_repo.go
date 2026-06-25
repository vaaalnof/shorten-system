package db

import (
	"context"
	"time"

	"shortener-service/internal/entity"
	"shortener-service/internal/repository"
	"shortener-service/internal/repository/db/query"
	"shortener-service/internal/usecase/port"
)

var _ port.ReportRepository = (*ReportRepo)(nil)

type ReportRepo struct {
	repo *repository.Repository
}

func NewReportRepo(
	repo *repository.Repository,
) *ReportRepo {

	return &ReportRepo{
		repo: repo,
	}
}

func (r *ReportRepo) GetSummary(
	ctx context.Context,
	urlID string,
) (
	*entity.ReportSummary,
	error,
) {

	now := time.Now().UTC()

	today := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		0,
		0,
		0,
		0,
		time.UTC,
	).Unix()

	summary := &entity.ReportSummary{}

	err := r.repo.QueryRow(
		ctx,
		query.GetReportSummary,
		urlID,
		today,
	).Scan(
		&summary.TotalClicks,
		&summary.UniqueVisitors,
		&summary.TodayClicks,
		&summary.TodayVisitors,
	)

	if err != nil {

		return nil, err
	}

	return summary, nil
}

func (r *ReportRepo) GetChart(
	ctx context.Context,
	urlID string,
) (
	[]*entity.ReportChart,
	error,
) {

	rows, err := r.repo.Query(
		ctx,
		query.GetReportChart,
		urlID,
	)

	if err != nil {

		return nil, err
	}

	defer rows.Close()

	var items []*entity.ReportChart

	for rows.Next() {

		item := &entity.ReportChart{}

		if err := rows.Scan(
			&item.AnalyticsDate,
			&item.TotalClicks,
			&item.UniqueVisitors,
		); err != nil {

			return nil, err
		}

		items = append(
			items,
			item,
		)
	}

	if err := rows.Err(); err != nil {

		return nil, err
	}

	return items, nil
}

func (r *ReportRepo) GetReferrers(
	ctx context.Context,
	urlID string,
) (
	[]*entity.ReportReferrer,
	error,
) {

	rows, err := r.repo.Query(
		ctx,
		query.GetReportReferrers,
		urlID,
	)

	if err != nil {

		return nil, err
	}

	defer rows.Close()

	var items []*entity.ReportReferrer

	for rows.Next() {

		item := &entity.ReportReferrer{}

		if err := rows.Scan(
			&item.Referrer,
			&item.Clicks,
		); err != nil {

			return nil, err
		}

		items = append(
			items,
			item,
		)
	}

	if err := rows.Err(); err != nil {

		return nil, err
	}

	return items, nil
}

func (r *ReportRepo) GetCountries(
	ctx context.Context,
	urlID string,
) (
	[]*entity.ReportCountry,
	error,
) {

	rows, err := r.repo.Query(
		ctx,
		query.GetReportCountries,
		urlID,
	)

	if err != nil {

		return nil, err
	}

	defer rows.Close()

	var items []*entity.ReportCountry

	for rows.Next() {

		item := &entity.ReportCountry{}

		if err := rows.Scan(
			&item.Country,
			&item.Clicks,
		); err != nil {

			return nil, err
		}

		items = append(
			items,
			item,
		)
	}

	if err := rows.Err(); err != nil {

		return nil, err
	}

	return items, nil
}

func (r *ReportRepo) GetDevices(
	ctx context.Context,
	urlID string,
) (
	[]*entity.ReportDevice,
	error,
) {

	rows, err := r.repo.Query(
		ctx,
		query.GetReportDevices,
		urlID,
	)

	if err != nil {

		return nil, err
	}

	defer rows.Close()

	var items []*entity.ReportDevice

	for rows.Next() {

		item := &entity.ReportDevice{}

		if err := rows.Scan(
			&item.Device,
			&item.Clicks,
		); err != nil {

			return nil, err
		}

		items = append(
			items,
			item,
		)
	}

	if err := rows.Err(); err != nil {

		return nil, err
	}

	return items, nil
}
