package port

import (
	"auth-service/internal/entity"
	"context"
)

type UserSessionRepository interface {
	Create(
		ctx context.Context,
		session *entity.UserSession,
	) error

	FindByID(
		ctx context.Context,
		id string,
	) (*entity.UserSession, error)

	FindValidByID(
		ctx context.Context,
		id string,
	) (*entity.UserSession, error)

	UpdateRefreshToken(
		ctx context.Context,
		sessionID string,
		refreshToken string,
	) error

	Revoke(
		ctx context.Context,
		id string,
		revokedAt int64,
	) error
}
