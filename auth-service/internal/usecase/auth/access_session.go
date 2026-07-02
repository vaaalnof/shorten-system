package auth

import (
	"context"
	"encoding/json"

	"auth-service/internal/delivery/http/middleware"
	"auth-service/internal/exception"
	"auth-service/internal/model"
	"auth-service/internal/security"
	"auth-service/internal/usecase/port"
	cachekey "auth-service/internal/utils/cache"
)

type AccessSessionUseCase struct {
	cache      port.Cache
	jwtService security.JWTService
}

func NewAccessSessionUseCase(
	cache port.Cache,
	jwtService security.JWTService,
) *AccessSessionUseCase {

	return &AccessSessionUseCase{
		cache:      cache,
		jwtService: jwtService,
	}
}

func (u *AccessSessionUseCase) GetSession(
	ctx context.Context,
) (*model.SessionCache, error) {

	// =====================================================
	// REQUEST META
	// =====================================================

	meta := middleware.GetMeta(
		ctx,
	)

	if meta == nil ||
		meta.Auth == "" {

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
	// TOKEN TYPE
	// =====================================================

	tokenType, ok := claims["typ"].(string)

	if !ok ||
		tokenType != "access" {

		return nil, exception.Unauthorized(
			"invalid token",
		)
	}

	// =====================================================
	// USER ID
	// =====================================================

	userID, ok := claims["sub"].(string)

	if !ok ||
		userID == "" {

		return nil, exception.Unauthorized(
			"invalid token",
		)
	}

	// =====================================================
	// SESSION ID
	// =====================================================

	sessionID, ok := claims["sid"].(string)

	if !ok ||
		sessionID == "" {

		return nil, exception.Unauthorized(
			"invalid token",
		)
	}

	// =====================================================
	// SESSION CACHE
	// =====================================================

	cacheValue, err := u.cache.Get(
		ctx,
		cachekey.Session(
			sessionID,
		),
	)

	if err != nil {

		return nil, exception.Internal(
			"failed to validate session",
		)
	}

	if cacheValue == "" {

		return nil, exception.Unauthorized(
			"session expired",
		)
	}

	// =====================================================
	// PARSE CACHE
	// =====================================================

	var session model.SessionCache

	if err := json.Unmarshal(
		[]byte(cacheValue),
		&session,
	); err != nil {

		return nil, exception.Internal(
			"failed to validate session",
		)
	}

	// =====================================================
	// RESPONSE
	// =====================================================

	return &model.SessionCache{
		UserID:        userID,
		SessionID:     sessionID,
		EmailVerified: session.EmailVerified,
	}, nil
}
