package auth

import (
	"auth-service/internal/delivery/http/middleware"
	"auth-service/internal/entity"
	"auth-service/internal/exception"
	"auth-service/internal/model"
	"auth-service/internal/model/converter"
	"auth-service/internal/repository"
	"auth-service/internal/security"
	"auth-service/internal/usecase/port"
	avatarutil "auth-service/internal/utils/avatar"
	cachekey "auth-service/internal/utils/cache"
	"time"

	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type RegisterUseCase struct {
	repo                 *repository.Repository
	userRepo             port.UserRepository
	validate             *validator.Validate
	userAuthProviderRepo port.UserAuthProviderRepository
	passwordHash         security.PasswordHash
	rateLimiter          security.RateLimiter
	registerMaxAttempts  int
	registerWindowTTL    time.Duration
}

func NewRegisterUseCase(
	repo *repository.Repository,
	validate *validator.Validate,
	userRepo port.UserRepository,
	userAuthProviderRepo port.UserAuthProviderRepository,
	passwordHash security.PasswordHash,
	rateLimiter security.RateLimiter,
	registerMaxAttempts int,
	registerWindowTTL time.Duration,
) *RegisterUseCase {
	return &RegisterUseCase{
		repo:                 repo,
		validate:             validate,
		userRepo:             userRepo,
		userAuthProviderRepo: userAuthProviderRepo,
		passwordHash:         passwordHash,
		rateLimiter:          rateLimiter,
		registerMaxAttempts:  registerMaxAttempts,
		registerWindowTTL:    registerWindowTTL,
	}
}

func (u *RegisterUseCase) Execute(
	ctx context.Context,
	req *model.RegisterUserRequest,
) (*model.WebResponse[*model.RegisterUserResponse], error) {

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

	if meta != nil &&
		meta.ClientIP != "" {

		ip = meta.ClientIP
	}

	// =====================================================
	// RATE LIMIT
	// =====================================================

	rateLimitKey := cachekey.RateLimitRegister(
		ip,
	)

	if err := u.rateLimiter.Check(
		ctx,
		rateLimitKey,
		u.registerMaxAttempts,
	); err != nil {

		return nil, err
	}

	// =====================================================
	// CHECK EMAIL
	// =====================================================

	existingUser, err := u.userRepo.FindByEmail(
		ctx,
		req.Email,
	)

	if err != nil {
		return nil, exception.Internal(
			"registration failed",
		)
	}

	if existingUser != nil {

		_ = u.rateLimiter.Increment(
			ctx,
			rateLimitKey,
			u.registerWindowTTL,
		)

		return nil, exception.Conflict(
			"email already registered",
		)
	}

	// =====================================================
	// HASH PASSWORD
	// =====================================================

	hashedPassword, err := u.passwordHash.Hash(
		req.Password,
	)

	if err != nil {
		return nil, exception.Internal(
			"registration failed",
		)
	}

	// =====================================================
	// PREPARE ENTITY
	// =====================================================

	lastName := ""

	if req.LastName != nil {
		lastName = *req.LastName
	}

	user := &entity.User{
		ID:        uuid.NewString(),
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  lastName,
		AvatarURL: avatarutil.DefaultAvatar(
			req.FirstName,
			lastName,
		),
		IsActive:      true,
		EmailVerified: false,
	}

	authProvider := &entity.UserAuthProvider{
		ID:           uuid.NewString(),
		UserID:       user.ID,
		Provider:     "local",
		PasswordHash: &hashedPassword,
	}

	// =====================================================
	// TRANSACTION
	// =====================================================

	err = u.repo.WithTransaction(
		ctx,
		func(
			ctx context.Context,
			tx *sql.Tx,
		) error {

			if err := u.userRepo.AddUser(
				ctx,
				tx,
				user,
			); err != nil {
				return err
			}

			if err := u.userAuthProviderRepo.AddAuthProvider(
				ctx,
				tx,
				authProvider,
			); err != nil {
				return err
			}

			return nil
		},
	)

	if err != nil {
		return nil, err
	}

	// =====================================================
	// REGISTER RATE LIMIT COUNTER
	// =====================================================

	_ = u.rateLimiter.Increment(
		ctx,
		rateLimitKey,
		u.registerWindowTTL,
	)

	// =====================================================
	// RESPONSE
	// =====================================================

	return &model.WebResponse[*model.RegisterUserResponse]{
		Message: "registration success",
		Data: converter.ToRegisterUserResponse(
			user,
		),
	}, nil
}
