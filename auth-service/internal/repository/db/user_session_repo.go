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

var _ port.UserSessionRepository = (*UserSessionRepo)(nil)

type UserSessionRepo struct {
	repo *repository.Repository
}

func NewUserSessionRepo(
	repo *repository.Repository,
) *UserSessionRepo {
	return &UserSessionRepo{
		repo: repo,
	}
}

func (r *UserSessionRepo) Create(
	ctx context.Context,
	session *entity.UserSession,
) error {

	session.CreatedAt = time.Now().Unix()

	_, err := r.repo.Exec(
		ctx,
		query.UserSessionCreate,
		session.ID,
		session.UserID,
		session.RefreshToken,
		session.IPAddress,
		session.UserAgent,
		session.ExpiredAt,
		session.RevokedAt,
		session.CreatedAt,
	)

	return err
}

func (r *UserSessionRepo) FindByID(
	ctx context.Context,
	id string,
) (*entity.UserSession, error) {

	row := r.repo.QueryRow(
		ctx,
		query.UserSessionFindByID,
		id,
	)

	return scanUserSession(
		row,
	)
}

func (r *UserSessionRepo) FindValidByID(
	ctx context.Context,
	id string,
) (*entity.UserSession, error) {

	row := r.repo.QueryRow(
		ctx,
		query.UserSessionFindValidByID,
		id,
		time.Now().Unix(),
	)

	return scanUserSession(
		row,
	)
}

func (r *UserSessionRepo) UpdateRefreshToken(
	ctx context.Context,
	sessionID string,
	refreshToken string,
) error {

	_, err := r.repo.Exec(
		ctx,
		query.UserSessionUpdateRefreshToken,
		refreshToken,
		sessionID,
	)

	return err
}

func (r *UserSessionRepo) Revoke(
	ctx context.Context,
	id string,
	revokedAt int64,
) error {

	_, err := r.repo.Exec(
		ctx,
		query.UserSessionRevoke,
		revokedAt,
		id,
	)

	return err
}

func scanUserSession(
	row interface {
		Scan(dest ...any) error
	},
) (*entity.UserSession, error) {

	session := &entity.UserSession{}

	err := row.Scan(
		&session.ID,
		&session.UserID,
		&session.RefreshToken,
		&session.IPAddress,
		&session.UserAgent,
		&session.ExpiredAt,
		&session.RevokedAt,
		&session.CreatedAt,
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

	return session, nil
}
