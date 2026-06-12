package usecase

import (
	"auth-service/internal/delivery/http/middleware"
	"auth-service/internal/exception"
	"auth-service/internal/model"
	"auth-service/internal/security"
	"auth-service/internal/usecase/port"
	cachekey "auth-service/internal/utils/cache"

	"context"
	"time"
)

type LogoutUseCase struct {
	userSessionRepo port.UserSessionRepository
	cache           port.Cache
	jwtService      security.JWTService
}

func NewLogoutUseCase(
	userSessionRepo port.UserSessionRepository,
	cache port.Cache,
	jwtService security.JWTService,
) *LogoutUseCase {
	return &LogoutUseCase{
		userSessionRepo: userSessionRepo,
		cache:           cache,
		jwtService:      jwtService,
	}
}

func (u *LogoutUseCase) Execute(
	ctx context.Context,
) (*model.WebResponse[any], error) {

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
	// SESSION ID
	// =====================================================

	sessionID, ok := claims["sid"].(string)

	if !ok || sessionID == "" {
		return nil, exception.Unauthorized(
			"invalid token",
		)
	}

	// =====================================================
	// SESSION
	// =====================================================

	session, err := u.userSessionRepo.FindValidByID(
		ctx,
		sessionID,
	)

	if err != nil {
		return nil, exception.Internal(
			"logout failed",
		)
	}

	if session == nil {
		return nil, exception.Unauthorized(
			"session not found",
		)
	}

	// =====================================================
	// REVOKE SESSION
	// =====================================================

	err = u.userSessionRepo.Revoke(
		ctx,
		sessionID,
		time.Now().Unix(),
	)

	if err != nil {
		return nil, exception.Internal(
			"logout failed",
		)
	}

	// =====================================================
	// DELETE CACHE
	// =====================================================

	_ = u.cache.Delete(
		ctx,
		cachekey.Session(
			sessionID,
		),
	)

	// =====================================================
	// RESPONSE
	// =====================================================

	return &model.WebResponse[any]{
		Message: "logout success",
		Data:    nil,
	}, nil
}
