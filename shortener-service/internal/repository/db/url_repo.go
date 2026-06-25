package db

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"shortener-service/internal/entity"
	"shortener-service/internal/repository"
	"shortener-service/internal/repository/db/query"
	"shortener-service/internal/usecase/port"
)

var _ port.URLRepository = (*URLRepo)(nil)

type URLRepo struct {
	repo *repository.Repository
}

func NewURLRepo(
	repo *repository.Repository,
) *URLRepo {

	return &URLRepo{
		repo: repo,
	}
}

func (r *URLRepo) Add(
	ctx context.Context,
	url *entity.URL,
) error {

	now := time.Now().Unix()

	url.CreatedAt = now
	url.UpdatedAt = now

	_, err := r.repo.Exec(
		ctx,
		query.AddURL,
		url.ID,
		url.UserID,
		url.ShortCode,
		url.OriginalURL,
		url.IsActive,
		url.PasswordHash,
		url.ExpiredAt,
		url.DeletedAt,
		url.CreatedAt,
		url.UpdatedAt,
	)

	return err
}

func (r *URLRepo) FindByShortCode(
	ctx context.Context,
	shortCode string,
) (*entity.URL, error) {

	row := r.repo.QueryRow(
		ctx,
		query.FindURLByShortCode,
		shortCode,
	)

	return scanURL(
		row,
	)
}

func (r *URLRepo) FindByID(
	ctx context.Context,
	id string,
) (*entity.URL, error) {

	row := r.repo.QueryRow(
		ctx,
		query.FindURLByID,
		id,
	)

	return scanURL(
		row,
	)
}

func (r *URLRepo) UpdatePassword(
	ctx context.Context,
	url *entity.URL,
) error {

	now := time.Now().Unix()

	url.UpdatedAt = now

	_, err := r.repo.Exec(
		ctx,
		query.UpdateURLPassword,
		url.ID,
		url.UserID,
		url.PasswordHash,
		url.UpdatedAt,
	)

	return err
}

func (r *URLRepo) RemovePassword(
	ctx context.Context,
	url *entity.URL,
) error {

	now := time.Now().Unix()

	url.PasswordHash = nil
	url.UpdatedAt = now

	_, err := r.repo.Exec(
		ctx,
		query.RemoveURLPassword,
		url.ID,
		url.UserID,
		url.UpdatedAt,
	)

	return err
}

func (r *URLRepo) CountByUserID(
	ctx context.Context,
	userID string,
) (
	int64,
	error,
) {

	var total int64

	err := r.repo.QueryRow(
		ctx,
		query.CountURLsByUserID,
		userID,
	).Scan(
		&total,
	)

	if err != nil {

		return 0, err
	}

	return total, nil
}

func (r *URLRepo) ListByUserID(
	ctx context.Context,
	userID string,
	limit int,
	offset int,
) (
	[]*entity.URL,
	error,
) {

	rows, err := r.repo.Query(
		ctx,
		query.ListURLsByUserID,
		userID,
		limit,
		offset,
	)

	if err != nil {

		return nil, err
	}

	defer rows.Close()

	return scanURLs(
		rows,
	)
}

func scanURLFields(
	scanner interface {
		Scan(dest ...any) error
	},
	url *entity.URL,
) error {

	return scanner.Scan(
		&url.ID,
		&url.UserID,
		&url.ShortCode,
		&url.OriginalURL,
		&url.IsActive,
		&url.PasswordHash,
		&url.ExpiredAt,
		&url.DeletedAt,
		&url.CreatedAt,
		&url.UpdatedAt,
	)
}

func scanURL(
	row interface {
		Scan(dest ...any) error
	},
) (*entity.URL, error) {

	url := &entity.URL{}

	err := scanURLFields(
		row,
		url,
	)

	if errors.Is(
		err,
		sql.ErrNoRows,
	) {

		return nil, nil
	}

	if err != nil {

		return nil, err
	}

	return url, nil
}

func scanURLs(
	rows interface {
		Next() bool
		Scan(dest ...any) error
		Err() error
	},
) (
	[]*entity.URL,
	error,
) {

	var urls []*entity.URL

	for rows.Next() {

		url := &entity.URL{}

		if err := scanURLFields(
			rows,
			url,
		); err != nil {

			return nil, err
		}

		urls = append(
			urls,
			url,
		)
	}

	if err := rows.Err(); err != nil {

		return nil, err
	}

	return urls, nil
}

func (r *URLRepo) Delete(
	ctx context.Context,
	url *entity.URL,
) error {

	now := time.Now().Unix()

	url.DeletedAt = &now
	url.UpdatedAt = now

	_, err := r.repo.Exec(
		ctx,
		query.DeleteURL,
		url.ID,
		url.UserID,
		url.DeletedAt,
		url.UpdatedAt,
	)

	return err
}

func (r *URLRepo) Update(
	ctx context.Context,
	url *entity.URL,
) error {

	now := time.Now().Unix()

	url.UpdatedAt = now

	_, err := r.repo.Exec(
		ctx,
		query.UpdateURL,
		url.ID,
		url.UserID,
		url.OriginalURL,
		url.ShortCode,
		url.IsActive,
		url.ExpiredAt,
		url.UpdatedAt,
	)

	return err
}
