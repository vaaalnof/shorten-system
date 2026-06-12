package usecase

import (
	"auth-service/internal/delivery/http/middleware"
	"auth-service/internal/exception"
	"auth-service/internal/model"
	"auth-service/internal/model/converter"
	"auth-service/internal/security"
	"auth-service/internal/usecase/port"
	cachekey "auth-service/internal/utils/cache"

	"context"
)

type ValidateTokenUseCase struct {
	cache      port.Cache
	jwtService security.JWTService
}

func NewValidateTokenUseCase(
	cache port.Cache,
	jwtService security.JWTService,
) *ValidateTokenUseCase {
	return &ValidateTokenUseCase{
		cache:      cache,
		jwtService: jwtService,
	}
}

func (u *ValidateTokenUseCase) Execute(
	ctx context.Context,
) (*model.WebResponse[*model.ValidateTokenResponse], error) {

	// =====================================================
	// REQUEST META
	// =====================================================

	meta := middleware.GetMeta(
		ctx,
	)

	if meta == nil {
		return nil, exception.Unauthorized(
			"missing authorization token",
		)
	}

	if meta.Auth == "" {
		return nil, exception.Unauthorized(
			"missing authorization token",
		)
	}

	// =====================================================
	// PARSE TOKEN
	// =====================================================

	claims, err := u.jwtService.ParseToken(
		meta.Auth,
	)

	if err != nil {
		return nil, exception.Unauthorized(
			"invalid token",
		)
	}

	// =====================================================
	// USER ID
	// =====================================================

	userID, ok := claims["sub"].(string)

	if !ok || userID == "" {
		return nil, exception.Unauthorized(
			"invalid token",
		)
	}

	// =====================================================
	// SESSION ID
	// =====================================================

	sessionID, ok := claims["sid"].(string)

	if !ok || sessionID == "" {
		return nil, exception.Unauthorized(
			"invalid token",
		)
	}

	// =====================================================
	// SESSION CHECK
	// =====================================================

	if !u.cache.Exists(
		ctx,
		cachekey.Session(
			sessionID,
		),
	) {
		return nil, exception.Unauthorized(
			"session expired",
		)
	}

	// =====================================================
	// RESPONSE
	// =====================================================

	return &model.WebResponse[*model.ValidateTokenResponse]{
		Message: "token valid",
		Data: converter.ToValidateTokenResponse(
			userID,
			sessionID,
		),
	}, nil
}
