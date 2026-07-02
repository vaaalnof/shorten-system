package auth

import (
	"auth-service/internal/entity"
	"context"
	"time"

	"auth-service/internal/delivery/http/middleware"
	"auth-service/internal/exception"
	"auth-service/internal/model"
	"auth-service/internal/model/converter"
	"auth-service/internal/security"
	"auth-service/internal/usecase/port"
	cachekey "auth-service/internal/utils/cache"

	"github.com/go-playground/validator/v10"
)

type LoginUseCase struct {
	validate             *validator.Validate
	userRepo             port.UserRepository
	userAuthProviderRepo port.UserAuthProviderRepository

	createSession *CreateSession

	passwordHash security.PasswordHash

	rateLimiter security.RateLimiter

	loginMaxAttempts int
	loginWindowTTL   time.Duration
}

func NewLoginUseCase(
	validate *validator.Validate,
	userRepo port.UserRepository,
	userAuthProviderRepo port.UserAuthProviderRepository,
	createSession *CreateSession,
	passwordHash security.PasswordHash,
	rateLimiter security.RateLimiter,
	loginMaxAttempts int,
	loginWindowTTL time.Duration,
) *LoginUseCase {

	return &LoginUseCase{
		validate: validate,

		userRepo: userRepo,

		userAuthProviderRepo: userAuthProviderRepo,

		createSession: createSession,

		passwordHash: passwordHash,

		rateLimiter: rateLimiter,

		loginMaxAttempts: loginMaxAttempts,
		loginWindowTTL:   loginWindowTTL,
	}
}

func (u *LoginUseCase) Execute(
	ctx context.Context,
	req *model.LoginRequest,
) (*model.WebResponse[*model.LoginResponse], error) {

	// =====================================================
	// VALIDATION
	// =====================================================

	if err := u.validate.Struct(
		req,
	); err != nil {

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

	if meta != nil &&
		meta.ClientIP != "" {

		ip = meta.ClientIP
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

	authProvider, err := u.userAuthProviderRepo.FindByEmailAndProvider(
		ctx,
		req.Email,
		string(entity.AuthProviderLocal),
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

	user, err := u.userRepo.FindByUserID(
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
	// CREATE SESSION
	// =====================================================

	loginResponse, err := u.createSession.Execute(
		ctx,
		user,
	)

	if err != nil {

		return nil, err
	}

	// =====================================================
	// RESET RATE LIMIT
	// =====================================================

	_ = u.rateLimiter.Reset(
		ctx,
		rateLimitKey,
	)

	// =====================================================
	// RESPONSE
	// =====================================================

	return &model.WebResponse[*model.LoginResponse]{
		Message: "login success",
		Data: converter.ToLoginResponse(
			loginResponse.AccessToken,
			loginResponse.RefreshToken,
		),
	}, nil
}
