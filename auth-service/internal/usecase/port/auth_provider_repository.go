package port

import (
	"context"
	"database/sql"

	"auth-service/internal/entity"
)

type UserAuthProviderRepository interface {
	CreateTx(
		ctx context.Context,
		tx *sql.Tx,
		auth *entity.UserAuthProvider,
	) error

	FindLocalByEmail(
		ctx context.Context,
		email string,
	) (*entity.UserAuthProvider, error)
}
