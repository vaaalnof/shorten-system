package auth

import (
	"context"
	"encoding/json"
	"time"

	"auth-service/internal/exception"
	"auth-service/internal/model"
	"auth-service/internal/model/converter"
	"auth-service/internal/usecase/port"
	cachekey "auth-service/internal/utils/cache"

	"github.com/google/uuid"
)

type GoogleLoginUseCase struct {
	cache port.Cache

	googleOAuth port.GoogleOAuthService

	oauthStateTTL time.Duration
}

func NewGoogleLoginUseCase(
	cache port.Cache,
	googleOAuth port.GoogleOAuthService,
	oauthStateTTL time.Duration,
) *GoogleLoginUseCase {

	return &GoogleLoginUseCase{
		cache: cache,

		googleOAuth: googleOAuth,

		oauthStateTTL: oauthStateTTL,
	}
}

func (u *GoogleLoginUseCase) Execute(
	ctx context.Context,
) (*model.WebResponse[*model.GoogleLoginResponse], error) {

	state := uuid.NewString()

	cacheValue, err := json.Marshal(
		&model.OAuthStateCache{
			Provider: "google",

			CreatedAt: time.Now().Unix(),
		},
	)

	if err != nil {

		return nil, exception.Internal(
			"failed to generate google login url",
		)
	}

	err = u.cache.Set(
		ctx,
		cachekey.OAuthState(
			"google",
			state,
		),
		string(cacheValue),
		u.oauthStateTTL,
	)

	if err != nil {

		return nil, exception.Internal(
			"failed to generate google login url",
		)
	}

	loginURL := u.googleOAuth.AuthURL(
		state,
	)

	return &model.WebResponse[*model.GoogleLoginResponse]{
		Message: "google login url generated",

		Data: converter.ToGoogleLoginResponse(
			loginURL,
		),
	}, nil
}
