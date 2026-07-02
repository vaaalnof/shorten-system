package auth

import (
	"auth-service/internal/model"
	"context"

	"auth-service/internal/delivery/http/middleware"
	"auth-service/internal/exception"
	"auth-service/internal/security"
)

type RefreshSessionUseCase struct {
	jwtService security.JWTService
}

func NewRefreshSessionUseCase(
	jwtService security.JWTService,
) *RefreshSessionUseCase {

	return &RefreshSessionUseCase{
		jwtService: jwtService,
	}
}

func (u *RefreshSessionUseCase) GetSession(
	ctx context.Context,
) (*model.SessionCache, error) {

	meta := middleware.GetMeta(
		ctx,
	)

	if meta == nil ||
		meta.RefreshToken == "" {

		return nil, exception.Unauthorized(
			"missing refresh token",
		)
	}

	claims, err := u.jwtService.ParseToken(
		meta.RefreshToken,
	)

	if err != nil {

		return nil, exception.Unauthorized(
			"invalid refresh token",
		)
	}

	tokenType, ok := claims["typ"].(string)

	if !ok ||
		tokenType != "refresh" {

		return nil, exception.Unauthorized(
			"invalid refresh token",
		)
	}

	userID, ok := claims["sub"].(string)

	if !ok ||
		userID == "" {

		return nil, exception.Unauthorized(
			"invalid refresh token",
		)
	}

	sessionID, ok := claims["sid"].(string)

	if !ok ||
		sessionID == "" {

		return nil, exception.Unauthorized(
			"invalid refresh token",
		)
	}

	return &model.SessionCache{
		UserID:       userID,
		SessionID:    sessionID,
		RefreshToken: meta.RefreshToken,
	}, nil
}
