package port

import (
	"context"
	"database/sql"

	"auth-service/internal/entity"
)

type UserRepository interface {
	CreateTx(
		ctx context.Context,
		tx *sql.Tx,
		user *entity.User,
	) error

	FindByEmail(
		ctx context.Context,
		email string,
	) (*entity.User, error)

	FindByID(
		ctx context.Context,
		id string,
	) (*entity.User, error)
}
