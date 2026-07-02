package port

import (
	"context"

	"auth-service/internal/model"
)

type GoogleOAuthService interface {
	AuthURL(
		state string,
	) string

	Exchange(
		ctx context.Context,
		code string,
	) (*model.GoogleUser, error)
}
