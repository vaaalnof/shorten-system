package auth

import (
	"context"
	"encoding/json"
	"time"

	"auth-service/internal/delivery/http/middleware"
	"auth-service/internal/entity"
	"auth-service/internal/exception"
	"auth-service/internal/model"
	"auth-service/internal/security"
	"auth-service/internal/usecase/port"
	cachekey "auth-service/internal/utils/cache"

	"github.com/google/uuid"
)

type CreateSession struct {
	userSessionRepo port.UserSessionRepository
	cache           port.Cache

	refreshTokenHash security.RefreshTokenHash
	jwtService       security.JWTService

	sessionTTL time.Duration
}

func NewCreateSession(
	userSessionRepo port.UserSessionRepository,
	cache port.Cache,
	refreshTokenHash security.RefreshTokenHash,
	jwtService security.JWTService,
	sessionTTL time.Duration,
) *CreateSession {

	return &CreateSession{
		userSessionRepo: userSessionRepo,
		cache:           cache,

		refreshTokenHash: refreshTokenHash,
		jwtService:       jwtService,

		sessionTTL: sessionTTL,
	}
}

func (u *CreateSession) Execute(
	ctx context.Context,
	user *entity.User,
) (*model.LoginResponse, error) {

	// =====================================================
	// REQUEST META
	// =====================================================

	meta := middleware.GetMeta(
		ctx,
	)

	var ipAddress *string
	var userAgent *string

	if meta != nil {

		if meta.ClientIP != "" {
			ipAddress = &meta.ClientIP
		}

		if meta.UserAgent != "" {
			userAgent = &meta.UserAgent
		}
	}

	// =====================================================
	// SESSION ID
	// =====================================================

	sessionID := uuid.NewString()

	// =====================================================
	// GENERATE ACCESS TOKEN
	// =====================================================

	accessToken, err := u.jwtService.GenerateAccessToken(
		user.ID,
		sessionID,
	)

	if err != nil {

		return nil, exception.Internal(
			"login failed",
		)
	}

	// =====================================================
	// GENERATE REFRESH TOKEN
	// =====================================================

	refreshToken, err := u.jwtService.GenerateRefreshToken(
		user.ID,
		sessionID,
	)

	if err != nil {

		return nil, exception.Internal(
			"login failed",
		)
	}

	// =====================================================
	// SAVE SESSION
	// =====================================================

	session := &entity.UserSession{
		ID:     sessionID,
		UserID: user.ID,

		RefreshToken: u.refreshTokenHash.Hash(
			refreshToken,
		),

		IPAddress: ipAddress,
		UserAgent: userAgent,

		ExpiredAt: time.Now().
			Add(
				u.jwtService.RefreshTokenTTL(),
			).
			Unix(),
	}

	if err := u.userSessionRepo.AddSession(
		ctx,
		session,
	); err != nil {

		return nil, exception.Internal(
			"login failed",
		)
	}

	// =====================================================
	// BUILD SESSION CACHE
	// =====================================================

	cacheValue, err := json.Marshal(
		&model.SessionCache{
			UserID:        user.ID,
			SessionID:     sessionID,
			EmailVerified: user.EmailVerified,
			ExpiredAt:     session.ExpiredAt,
		},
	)

	if err != nil {

		return nil, exception.Internal(
			"login failed",
		)
	}

	// =====================================================
	// CACHE SESSION
	// =====================================================

	if err := u.cache.Set(
		ctx,
		cachekey.Session(
			sessionID,
		),
		string(cacheValue),
		u.sessionTTL,
	); err != nil {

		return nil, exception.Internal(
			"login failed",
		)
	}

	// =====================================================
	// RESPONSE
	// =====================================================

	return &model.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
