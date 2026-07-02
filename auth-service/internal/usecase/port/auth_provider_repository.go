package port

import (
	"context"
	"database/sql"

	"auth-service/internal/entity"
)

type UserAuthProviderRepository interface {
	AddAuthProvider(
		ctx context.Context,
		tx *sql.Tx,
		auth *entity.UserAuthProvider,
	) error

	FindByEmailAndProvider(
		ctx context.Context,
		email string,
		provider string,
	) (*entity.UserAuthProvider, error)

	FindByProviderUserID(
		ctx context.Context,
		provider string,
		providerUserID string,
	) (*entity.UserAuthProvider, error)

	FindByUserIDAndProvider(
		ctx context.Context,
		userID string,
		provider string,
	) (*entity.UserAuthProvider, error)
}
