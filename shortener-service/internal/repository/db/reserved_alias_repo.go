package db

import (
	"context"

	"shortener-service/internal/repository"
	"shortener-service/internal/repository/db/query"
	"shortener-service/internal/usecase/port"
)

var _ port.ReservedAliasRepository = (*ReservedAliasRepo)(nil)

type ReservedAliasRepo struct {
	repo *repository.Repository
}

func NewReservedAliasRepo(
	repo *repository.Repository,
) *ReservedAliasRepo {

	return &ReservedAliasRepo{
		repo: repo,
	}
}

func (r *ReservedAliasRepo) Exists(
	ctx context.Context,
	keyword string,
) (bool, error) {

	row := r.repo.QueryRow(
		ctx,
		query.ReservedAliasExists,
		keyword,
	)

	var exists bool

	err := row.Scan(
		&exists,
	)

	if err != nil {
		return false, err
	}

	return exists, nil
}
