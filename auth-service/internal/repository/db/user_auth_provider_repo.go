package db

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"auth-service/internal/entity"
	"auth-service/internal/repository"
	"auth-service/internal/repository/db/query"
	"auth-service/internal/usecase/port"
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

func (r *UserAuthProviderRepo) AddAuthProvider(
	ctx context.Context,
	tx *sql.Tx,
	auth *entity.UserAuthProvider,
) error {

	auth.CreatedAt = time.Now().Unix()

	_, err := tx.ExecContext(
		ctx,
		query.AddAuthProvider,
		auth.ID,
		auth.UserID,
		auth.Provider,
		auth.ProviderUserID,
		auth.PasswordHash,
		auth.CreatedAt,
	)

	return err
}

func (r *UserAuthProviderRepo) FindByEmailAndProvider(
	ctx context.Context,
	email string,
	provider string,
) (*entity.UserAuthProvider, error) {

	row := r.repo.QueryRow(
		ctx,
		query.FindByEmailAndProvider,
		email,
		provider,
	)

	return scanUserAuthProvider(
		row,
	)
}

func (r *UserAuthProviderRepo) FindByProviderUserID(
	ctx context.Context,
	provider string,
	providerUserID string,
) (*entity.UserAuthProvider, error) {

	row := r.repo.QueryRow(
		ctx,
		query.FindByProviderUserID,
		provider,
		providerUserID,
	)

	return scanUserAuthProvider(
		row,
	)
}

func (r *UserAuthProviderRepo) FindByUserIDAndProvider(
	ctx context.Context,
	userID string,
	provider string,
) (*entity.UserAuthProvider, error) {

	row := r.repo.QueryRow(
		ctx,
		query.FindByUserIDAndProvider,
		userID,
		provider,
	)

	return scanUserAuthProvider(
		row,
	)
}

func scanUserAuthProvider(
	row interface {
		Scan(dest ...any) error
	},
) (*entity.UserAuthProvider, error) {

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
