package usecase

import (
	"auth-service/internal/model/converter"
	cachekey "auth-service/internal/utils/cache"
	"time"

	"auth-service/internal/delivery/http/middleware"
	"auth-service/internal/entity"
	"auth-service/internal/exception"
	"auth-service/internal/model"
	"auth-service/internal/security"
	"auth-service/internal/usecase/port"
	"context"
	"encoding/json"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type LoginUseCase struct {
	validate             *validator.Validate
	userRepo             port.UserRepository
	userAuthProviderRepo port.UserAuthProviderRepository
	userSessionRepo      port.UserSessionRepository
	cache                port.Cache
	passwordHash         security.PasswordHash
	refreshTokenHash     security.RefreshTokenHash
	jwtService           security.JWTService
	rateLimiter          security.RateLimiter
	loginMaxAttempts     int
	loginWindowTTL       time.Duration
	sessionTTL           time.Duration
}

func NewLoginUseCase(
	validate *validator.Validate,
	userRepo port.UserRepository,
	userAuthProviderRepo port.UserAuthProviderRepository,
	userSessionRepo port.UserSessionRepository,
	cache port.Cache,
	passwordHash security.PasswordHash,
	refreshTokenHash security.RefreshTokenHash,
	jwtService security.JWTService,
	rateLimiter security.RateLimiter,
	loginMaxAttempts int,
	loginWindowTTL time.Duration,
	sessionTTL time.Duration,
) *LoginUseCase {
	return &LoginUseCase{
		validate:             validate,
		userRepo:             userRepo,
		userAuthProviderRepo: userAuthProviderRepo,
		userSessionRepo:      userSessionRepo,
		cache:                cache,
		passwordHash:         passwordHash,
		refreshTokenHash:     refreshTokenHash,
		jwtService:           jwtService,
		rateLimiter:          rateLimiter,
		loginMaxAttempts:     loginMaxAttempts,
		loginWindowTTL:       loginWindowTTL,
		sessionTTL:           sessionTTL,
	}
}

func (u *LoginUseCase) Execute(
	ctx context.Context,
	req *model.LoginRequest,
) (*model.WebResponse[*model.LoginResponse], error) {

	// =====================================================
	// VALIDATION
	// =====================================================

	if err := u.validate.Struct(req); err != nil {
		return nil, exception.Validation(
			err,
		)
	}

	// =====================================================
	// REQUEST META
	// =====================================================

	meta := middleware.GetMeta(
		ctx,
	)

	ip := "unknown"

	var ipAddress *string
	var userAgent *string

	if meta != nil {

		if meta.ClientIP != "" {
			ip = meta.ClientIP
			ipAddress = &meta.ClientIP
		}

		if meta.UserAgent != "" {
			userAgent = &meta.UserAgent
		}
	}

	// =====================================================
	// RATE LIMIT
	// =====================================================

	rateLimitKey := cachekey.RateLimitLogin(
		ip,
		req.Email,
	)

	if err := u.rateLimiter.Check(
		ctx,
		rateLimitKey,
		u.loginMaxAttempts,
	); err != nil {

		return nil, err
	}

	// =====================================================
	// FIND AUTH PROVIDER
	// =====================================================

	authProvider, err := u.userAuthProviderRepo.FindLocalByEmail(
		ctx,
		req.Email,
	)

	if err != nil {
		return nil, exception.Internal(
			"login failed",
		)
	}

	if authProvider == nil {

		_ = u.rateLimiter.Increment(
			ctx,
			rateLimitKey,
			u.loginWindowTTL,
		)

		return nil, exception.Unauthorized(
			"invalid email or password",
		)
	}

	// =====================================================
	// VERIFY PASSWORD
	// =====================================================

	if authProvider.PasswordHash == nil {

		_ = u.rateLimiter.Increment(
			ctx,
			rateLimitKey,
			u.loginWindowTTL,
		)

		return nil, exception.Unauthorized(
			"invalid email or password",
		)
	}

	if err := u.passwordHash.Compare(
		*authProvider.PasswordHash,
		req.Password,
	); err != nil {

		_ = u.rateLimiter.Increment(
			ctx,
			rateLimitKey,
			u.loginWindowTTL,
		)

		return nil, exception.Unauthorized(
			"invalid email or password",
		)
	}

	// =====================================================
	// FIND USER
	// =====================================================

	user, err := u.userRepo.FindByID(
		ctx,
		authProvider.UserID,
	)

	if err != nil {
		return nil, exception.Internal(
			"login failed",
		)
	}

	if user == nil {

		_ = u.rateLimiter.Increment(
			ctx,
			rateLimitKey,
			u.loginWindowTTL,
		)

		return nil, exception.Unauthorized(
			"invalid email or password",
		)
	}

	// =====================================================
	// USER STATUS
	// =====================================================

	if !user.IsActive {

		_ = u.rateLimiter.Increment(
			ctx,
			rateLimitKey,
			u.loginWindowTTL,
		)

		return nil, exception.Unauthorized(
			"account is inactive",
		)
	}

	// =====================================================
	// SESSION ID
	// =====================================================

	sessionID := uuid.NewString()

	// =====================================================
	// GENERATE TOKENS
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

	if err := u.userSessionRepo.Create(
		ctx,
		session,
	); err != nil {

		return nil, exception.Internal(
			"login failed",
		)
	}

	// =====================================================
	// CACHE SESSION
	// =====================================================

	cacheValue, err := json.Marshal(
		&model.SessionCache{
			UserID:    user.ID,
			SessionID: sessionID,
			ExpiredAt: session.ExpiredAt,
		},
	)

	if err != nil {
		return nil, exception.Internal(
			"login failed",
		)
	}

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
	// RESET LOGIN RATE LIMIT
	// =====================================================

	_ = u.rateLimiter.Reset(
		ctx,
		rateLimitKey,
	)

	return &model.WebResponse[*model.LoginResponse]{
		Message: "login success",
		Data: converter.ToLoginResponse(
			accessToken,
			refreshToken,
		),
	}, nil
}
