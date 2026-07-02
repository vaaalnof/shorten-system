package port

import (
	"auth-service/internal/entity"
	"context"
	"database/sql"
)

type UserRepository interface {
	AddUser(
		ctx context.Context,
		tx *sql.Tx,
		user *entity.User,
	) error

	FindByEmail(
		ctx context.Context,
		email string,
	) (*entity.User, error)

	FindByUserID(
		ctx context.Context,
		id string,
	) (*entity.User, error)

	UpdateEmailVerified(
		ctx context.Context,
		userID string,
	) (int64, error)
}
