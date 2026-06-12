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
	"encoding/json"
	"time"
)

type MeUseCase struct {
	userRepo        port.UserRepository
	userSessionRepo port.UserSessionRepository
	cache           port.Cache
	jwtService      security.JWTService
	sessionTTL      time.Duration
}

func NewMeUseCase(
	userRepo port.UserRepository,
	userSessionRepo port.UserSessionRepository,
	cache port.Cache,
	jwtService security.JWTService,
	sessionTTL time.Duration,
) *MeUseCase {
	return &MeUseCase{
		userRepo:        userRepo,
		userSessionRepo: userSessionRepo,
		cache:           cache,
		jwtService:      jwtService,
		sessionTTL:      sessionTTL,
	}
}

func (u *MeUseCase) Execute(
	ctx context.Context,
) (*model.WebResponse[*model.MeResponse], error) {

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
	// SESSION CACHE
	// =====================================================

	sessionKey := cachekey.Session(
		sessionID,
	)

	sessionCache, err := u.cache.Get(
		ctx,
		sessionKey,
	)

	if err != nil {
		return nil, exception.Internal(
			"failed to get profile",
		)
	}

	// =====================================================
	// FALLBACK DATABASE
	// =====================================================

	if sessionCache == "" {

		session, err := u.userSessionRepo.FindValidByID(
			ctx,
			sessionID,
		)

		if err != nil {
			return nil, exception.Internal(
				"failed to get profile",
			)
		}

		if session == nil {
			return nil, exception.Unauthorized(
				"session expired",
			)
		}

		cacheValue, err := json.Marshal(
			&model.SessionCache{
				UserID:    session.UserID,
				SessionID: session.ID,
				ExpiredAt: session.ExpiredAt,
			},
		)

		if err == nil {

			_ = u.cache.Set(
				ctx,
				sessionKey,
				string(cacheValue),
				u.sessionTTL,
			)
		}
	}

	// =====================================================
	// USER
	// =====================================================

	user, err := u.userRepo.FindByID(
		ctx,
		userID,
	)

	if err != nil {
		return nil, exception.Internal(
			"failed to get profile",
		)
	}

	if user == nil {
		return nil, exception.Unauthorized(
			"user not found",
		)
	}

	// =====================================================
	// USER STATUS
	// =====================================================

	if !user.IsActive {
		return nil, exception.Unauthorized(
			"account is inactive",
		)
	}

	// =====================================================
	// RESPONSE
	// =====================================================

	return &model.WebResponse[*model.MeResponse]{
		Message: "profile fetched successfully",
		Data: converter.ToMeResponse(
			user,
		),
	}, nil
}
