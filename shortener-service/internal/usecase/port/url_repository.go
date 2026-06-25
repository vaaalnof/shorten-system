package port

import (
	"context"
	"shortener-service/internal/entity"
)

type URLRepository interface {
	Add(
		ctx context.Context,
		url *entity.URL,
	) error

	FindByShortCode(
		ctx context.Context,
		shortCode string,
	) (*entity.URL, error)

	FindByID(
		ctx context.Context,
		id string,
	) (*entity.URL, error)

	UpdatePassword(
		ctx context.Context,
		url *entity.URL,
	) error

	RemovePassword(
		ctx context.Context,
		url *entity.URL,
	) error

	CountByUserID(
		ctx context.Context,
		userID string,
	) (
		int64,
		error,
	)

	ListByUserID(
		ctx context.Context,
		userID string,
		limit int,
		offset int,
	) (
		[]*entity.URL,
		error,
	)

	Delete(
		ctx context.Context,
		url *entity.URL,
	) error

	Update(
		ctx context.Context,
		url *entity.URL,
	) error
}
