package user

import (
	"context"
	"encoding/json"
	"time"

	"auth-service/internal/delivery/http/middleware"
	"auth-service/internal/exception"
	"auth-service/internal/model"
	"auth-service/internal/security"
	"auth-service/internal/usecase/port"
	cachekey "auth-service/internal/utils/cache"
)

type UseCase struct {
	cache           port.Cache
	userSessionRepo port.UserSessionRepository
	jwtService      security.JWTService
	sessionTTL      time.Duration
}

type AuthUser struct {
	UserID    string
	SessionID string
}

func NewUseCase(
	cache port.Cache,
	userSessionRepo port.UserSessionRepository,
	jwtService security.JWTService,
	sessionTTL time.Duration,
) *UseCase {

	return &UseCase{
		cache:           cache,
		userSessionRepo: userSessionRepo,
		jwtService:      jwtService,
		sessionTTL:      sessionTTL,
	}
}

func (u *UseCase) GetAuthenticatedUser(
	ctx context.Context,
) (*AuthUser, error) {

	meta := middleware.GetMeta(
		ctx,
	)

	if meta == nil ||
		meta.Auth == "" {

		return nil, exception.Unauthorized(
			"missing authorization token",
		)
	}

	claims, err := u.jwtService.ParseToken(
		meta.Auth,
	)

	if err != nil {

		return nil, exception.Unauthorized(
			"invalid token",
		)
	}

	userID, ok := claims["sub"].(string)

	if !ok || userID == "" {

		return nil, exception.Unauthorized(
			"invalid token",
		)
	}

	sessionID, ok := claims["sid"].(string)

	if !ok || sessionID == "" {

		return nil, exception.Unauthorized(
			"invalid token",
		)
	}

	sessionKey := cachekey.Session(
		sessionID,
	)

	sessionCache, err := u.cache.Get(
		ctx,
		sessionKey,
	)

	if err != nil {

		return nil, exception.Internal(
			"failed to validate session",
		)
	}

	if sessionCache == "" {

		session, err := u.userSessionRepo.FindSessionByID(
			ctx,
			sessionID,
		)

		if err != nil {

			return nil, exception.Internal(
				"failed to validate session",
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

	return &AuthUser{
		UserID:    userID,
		SessionID: sessionID,
	}, nil
}
