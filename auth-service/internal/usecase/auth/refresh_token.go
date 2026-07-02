package auth

import (
	"context"
	"encoding/json"
	"time"

	"auth-service/internal/exception"
	"auth-service/internal/model"
	"auth-service/internal/model/converter"
	"auth-service/internal/security"
	"auth-service/internal/usecase/port"
	cachekey "auth-service/internal/utils/cache"
)

type RefreshTokenUseCase struct {
	*RefreshSessionUseCase

	userSessionRepo  port.UserSessionRepository
	cache            port.Cache
	refreshTokenHash security.RefreshTokenHash
	sessionTTL       time.Duration
}

func NewRefreshTokenUseCase(
	base *RefreshSessionUseCase,
	userSessionRepo port.UserSessionRepository,
	cache port.Cache,
	refreshTokenHash security.RefreshTokenHash,
	sessionTTL time.Duration,
) *RefreshTokenUseCase {

	return &RefreshTokenUseCase{
		RefreshSessionUseCase: base,
		userSessionRepo:       userSessionRepo,
		cache:                 cache,
		refreshTokenHash:      refreshTokenHash,
		sessionTTL:            sessionTTL,
	}
}

func (u *RefreshTokenUseCase) Execute(
	ctx context.Context,
) (*model.WebResponse[*model.RefreshTokenResponse], error) {

	// =====================================================
	// REFRESH SESSION
	// =====================================================

	refresh, err := u.GetSession(
		ctx,
	)

	if err != nil {
		return nil, err
	}

	// =====================================================
	// FIND SESSION
	// =====================================================

	session, err := u.userSessionRepo.FindSessionByID(
		ctx,
		refresh.SessionID,
	)

	if err != nil {

		return nil, exception.Internal(
			"refresh token failed",
		)
	}

	if session == nil {

		return nil, exception.Unauthorized(
			"invalid refresh token",
		)
	}

	// =====================================================
	// VERIFY REFRESH TOKEN
	// =====================================================

	hashedRefreshToken := u.refreshTokenHash.Hash(
		refresh.RefreshToken,
	)

	if session.RefreshToken != hashedRefreshToken {

		return nil, exception.Unauthorized(
			"invalid refresh token",
		)
	}

	// =====================================================
	// GENERATE TOKENS
	// =====================================================

	newAccessToken, err := u.jwtService.GenerateAccessToken(
		refresh.UserID,
		refresh.SessionID,
	)

	if err != nil {

		return nil, exception.Internal(
			"refresh token failed",
		)
	}

	newRefreshToken, err := u.jwtService.GenerateRefreshToken(
		refresh.UserID,
		refresh.SessionID,
	)

	if err != nil {

		return nil, exception.Internal(
			"refresh token failed",
		)
	}

	// =====================================================
	// UPDATE REFRESH TOKEN
	// =====================================================

	err = u.userSessionRepo.UpdateRefreshToken(
		ctx,
		refresh.SessionID,
		u.refreshTokenHash.Hash(
			newRefreshToken,
		),
	)

	if err != nil {

		return nil, exception.Internal(
			"refresh token failed",
		)
	}

	// =====================================================
	// REBUILD CACHE
	// =====================================================

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
			cachekey.Session(
				session.ID,
			),
			string(cacheValue),
			u.sessionTTL,
		)
	}

	// =====================================================
	// RESPONSE
	// =====================================================

	return &model.WebResponse[*model.RefreshTokenResponse]{
		Message: "token refreshed successfully",
		Data: converter.ToRefreshTokenResponse(
			newAccessToken,
			newRefreshToken,
		),
	}, nil
}
