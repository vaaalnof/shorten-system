package port

import (
	"auth-service/internal/entity"
	"context"
)

type UserSessionRepository interface {
	AddSession(
		ctx context.Context,
		session *entity.UserSession,
	) error

	FindSessionByID(
		ctx context.Context,
		id string,
	) (*entity.UserSession, error)

	FindSessionByUserID(
		ctx context.Context,
		userID string,
	) ([]*entity.UserSession, error)

	UpdateRefreshToken(
		ctx context.Context,
		sessionID string,
		refreshToken string,
	) error

	RevokeByID(
		ctx context.Context,
		id string,
		revokedAt int64,
	) error

	RevokeByUserID(
		ctx context.Context,
		userID string,
		revokedAt int64,
	) error
}
