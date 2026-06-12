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

type RefreshTokenUseCase struct {
	userSessionRepo  port.UserSessionRepository
	cache            port.Cache
	jwtService       security.JWTService
	refreshTokenHash security.RefreshTokenHash
	sessionTTL       time.Duration
}

func NewRefreshTokenUseCase(
	userSessionRepo port.UserSessionRepository,
	cache port.Cache,
	jwtService security.JWTService,
	refreshTokenHash security.RefreshTokenHash,
	sessionTTL time.Duration,
) *RefreshTokenUseCase {
	return &RefreshTokenUseCase{
		userSessionRepo:  userSessionRepo,
		cache:            cache,
		jwtService:       jwtService,
		refreshTokenHash: refreshTokenHash,
		sessionTTL:       sessionTTL,
	}
}

func (u *RefreshTokenUseCase) Execute(
	ctx context.Context,
) (*model.WebResponse[*model.RefreshTokenResponse], error) {

	// =====================================================
	// REQUEST META
	// =====================================================

	meta := middleware.GetMeta(
		ctx,
	)

	if meta == nil ||
		meta.RefreshToken == "" {

		return nil, exception.Unauthorized(
			"missing refresh token",
		)
	}

	refreshToken := meta.RefreshToken

	// =====================================================
	// PARSE REFRESH TOKEN
	// =====================================================

	claims, err := u.jwtService.ParseToken(
		refreshToken,
	)

	if err != nil {
		return nil, exception.Unauthorized(
			"invalid refresh token",
		)
	}

	// =====================================================
	// USER ID
	// =====================================================

	userID, ok := claims["sub"].(string)

	if !ok || userID == "" {
		return nil, exception.Unauthorized(
			"invalid refresh token",
		)
	}

	// =====================================================
	// SESSION ID
	// =====================================================

	sessionID, ok := claims["sid"].(string)

	if !ok || sessionID == "" {
		return nil, exception.Unauthorized(
			"invalid refresh token",
		)
	}

	// =====================================================
	// FIND SESSION
	// =====================================================

	session, err := u.userSessionRepo.FindValidByID(
		ctx,
		sessionID,
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
		refreshToken,
	)

	if session.RefreshToken != hashedRefreshToken {
		return nil, exception.Unauthorized(
			"invalid refresh token",
		)
	}

	// =====================================================
	// GENERATE ACCESS TOKEN
	// =====================================================

	newAccessToken, err := u.jwtService.GenerateAccessToken(
		userID,
		sessionID,
	)

	if err != nil {
		return nil, exception.Internal(
			"refresh token failed",
		)
	}

	// =====================================================
	// GENERATE REFRESH TOKEN
	// =====================================================

	newRefreshToken, err := u.jwtService.GenerateRefreshToken(
		userID,
		sessionID,
	)

	if err != nil {
		return nil, exception.Internal(
			"refresh token failed",
		)
	}

	// =====================================================
	// HASH NEW REFRESH TOKEN
	// =====================================================

	hashedNewRefreshToken := u.refreshTokenHash.Hash(
		newRefreshToken,
	)

	// =====================================================
	// UPDATE SESSION
	// =====================================================

	err = u.userSessionRepo.UpdateRefreshToken(
		ctx,
		sessionID,
		hashedNewRefreshToken,
	)

	if err != nil {
		return nil, exception.Internal(
			"refresh token failed",
		)
	}

	// =====================================================
	// BUILD SESSION CACHE
	// =====================================================

	cacheValue, err := json.Marshal(
		&model.SessionCache{
			UserID:    session.UserID,
			SessionID: session.ID,
			ExpiredAt: session.ExpiredAt,
		},
	)

	if err != nil {
		return nil, exception.Internal(
			"refresh token failed",
		)
	}

	// =====================================================
	// REBUILD SESSION CACHE
	// =====================================================

	err = u.cache.Set(
		ctx,
		cachekey.Session(
			session.ID,
		),
		string(cacheValue),
		u.sessionTTL,
	)

	if err != nil {
		return nil, exception.Internal(
			"refresh token failed",
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
