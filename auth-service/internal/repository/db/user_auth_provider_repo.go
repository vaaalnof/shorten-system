package db

import (
	"auth-service/internal/entity"
	"auth-service/internal/repository"
	"auth-service/internal/repository/db/query"
	"auth-service/internal/usecase/port"
	"context"
	"database/sql"
	"errors"
	"time"
)

var _ port.UserAuthProviderRepository = (*UserAuthProviderRepo)(nil)

type UserAuthProviderRepo struct {
	repo *repository.Repository
}

func NewUserAuthProviderRepo(
	repo *repository.Repository,
) *UserAuthProviderRepo {
	return &UserAuthProviderRepo{
		repo: repo,
	}
}

func (r *UserAuthProviderRepo) CreateTx(
	ctx context.Context,
	tx *sql.Tx,
	auth *entity.UserAuthProvider,
) error {

	auth.CreatedAt = time.Now().Unix()

	_, err := tx.ExecContext(
		ctx,
		query.UserAuthProviderCreate,
		auth.ID,
		auth.UserID,
		auth.Provider,
		auth.ProviderUserID,
		auth.PasswordHash,
		auth.CreatedAt,
	)

	return err
}

func (r *UserAuthProviderRepo) FindLocalByEmail(
	ctx context.Context,
	email string,
) (*entity.UserAuthProvider, error) {

	row := r.repo.QueryRow(
		ctx,
		query.UserAuthProviderFindLocalByEmail,
		email,
	)

	auth := &entity.UserAuthProvider{}

	err := row.Scan(
		&auth.ID,
		&auth.UserID,
		&auth.Provider,
		&auth.ProviderUserID,
		&auth.PasswordHash,
		&auth.CreatedAt,
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

	return auth, nil
}
