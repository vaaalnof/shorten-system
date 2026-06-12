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

var _ port.UserRepository = (*UserRepo)(nil)

type UserRepo struct {
	repo *repository.Repository
}

func NewUserRepo(
	repo *repository.Repository,
) *UserRepo {
	return &UserRepo{
		repo: repo,
	}
}

func (r *UserRepo) CreateTx(
	ctx context.Context,
	tx *sql.Tx,
	user *entity.User,
) error {

	now := time.Now().Unix()

	user.CreatedAt = now
	user.UpdatedAt = &now

	_, err := tx.ExecContext(
		ctx,
		query.UserCreate,
		user.ID,
		user.Email,
		user.FirstName,
		user.LastName,
		user.AvatarURL,
		user.IsActive,
		user.EmailVerified,
		user.CreatedAt,
		user.UpdatedAt,
	)

	return err
}

func (r *UserRepo) FindByEmail(
	ctx context.Context,
	email string,
) (*entity.User, error) {

	row := r.repo.QueryRow(
		ctx,
		query.UserFindByEmail,
		email,
	)

	user := &entity.User{}

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.AvatarURL,
		&user.IsActive,
		&user.EmailVerified,
		&user.CreatedAt,
		&user.UpdatedAt,
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

	return user, nil
}

func (r *UserRepo) FindByID(
	ctx context.Context,
	id string,
) (*entity.User, error) {

	row := r.repo.QueryRow(
		ctx,
		query.UserFindByID,
		id,
	)

	user := &entity.User{}

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.AvatarURL,
		&user.IsActive,
		&user.EmailVerified,
		&user.CreatedAt,
		&user.UpdatedAt,
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

	return user, nil
}
