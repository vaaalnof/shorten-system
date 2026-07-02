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

func (r *UserSessionRepo) AddSession(
	ctx context.Context,
	session *entity.UserSession,
) error {

	now := time.Now().Unix()

	session.CreatedAt = now
	session.LastSeenAt = now

	_, err := r.repo.Exec(
		ctx,
		query.AddSession,
		session.ID,
		session.UserID,
		session.RefreshToken,
		session.IPAddress,
		session.UserAgent,
		session.LastSeenAt,
		session.ExpiredAt,
		session.RevokedAt,
		session.CreatedAt,
	)

	return err
}

func (r *UserSessionRepo) FindSessionByID(
	ctx context.Context,
	id string,
) (*entity.UserSession, error) {

	row := r.repo.QueryRow(
		ctx,
		query.FindSessionByID,
		id,
		time.Now().Unix(),
	)

	return scanUserSession(
		row,
	)
}

func (r *UserSessionRepo) FindSessionByUserID(
	ctx context.Context,
	userID string,
) ([]*entity.UserSession, error) {

	rows, err := r.repo.Query(
		ctx,
		query.FindSessionByUserID,
		userID,
		time.Now().Unix(),
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var sessions []*entity.UserSession

	for rows.Next() {

		session := &entity.UserSession{}

		err := rows.Scan(
			&session.ID,
			&session.UserID,
			&session.RefreshToken,
			&session.IPAddress,
			&session.UserAgent,
			&session.ExpiredAt,
			&session.RevokedAt,
			&session.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		sessions = append(
			sessions,
			session,
		)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return sessions, nil
}

func (r *UserSessionRepo) UpdateRefreshToken(
	ctx context.Context,
	sessionID string,
	refreshToken string,
) error {

	_, err := r.repo.Exec(
		ctx,
		query.UpdateRefreshToken,
		refreshToken,
		time.Now().Unix(),
		sessionID,
	)

	return err
}

func (r *UserSessionRepo) RevokeByID(
	ctx context.Context,
	id string,
	revokedAt int64,
) error {

	_, err := r.repo.Exec(
		ctx,
		query.RevokeByID,
		revokedAt,
		id,
	)

	return err
}

func (r *UserSessionRepo) RevokeByUserID(
	ctx context.Context,
	userID string,
	revokedAt int64,
) error {

	_, err := r.repo.Exec(
		ctx,
		query.RevokeByUserID,
		revokedAt,
		userID,
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
