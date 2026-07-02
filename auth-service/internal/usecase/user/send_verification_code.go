package user

import (
	"auth-service/internal/exception"
	"auth-service/internal/model"
	"auth-service/internal/template"
	"auth-service/internal/usecase/port"
	cachekey "auth-service/internal/utils/cache"

	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

type SendVerificationCodeUseCase struct {
	*UseCase

	userRepo port.UserRepository
	mailer   port.Mailer

	emailVerificationTTL         time.Duration
	emailVerificationCooldownTTL time.Duration
}

func NewSendVerificationCodeUseCase(
	base *UseCase,
	userRepo port.UserRepository,
	mailer port.Mailer,
	emailVerificationTTL time.Duration,
	emailVerificationCooldownTTL time.Duration,
) *SendVerificationCodeUseCase {

	return &SendVerificationCodeUseCase{
		UseCase: base,

		userRepo: userRepo,
		mailer:   mailer,

		emailVerificationTTL:         emailVerificationTTL,
		emailVerificationCooldownTTL: emailVerificationCooldownTTL,
	}
}

func (u *SendVerificationCodeUseCase) Execute(
	ctx context.Context,
) (*model.WebResponse[any], error) {

	// =====================================================
	// AUTH USER
	// =====================================================

	authUser, err := u.GetAuthenticatedUser(
		ctx,
	)

	if err != nil {
		return nil, err
	}

	// =====================================================
	// USER
	// =====================================================

	user, err := u.userRepo.FindByUserID(
		ctx,
		authUser.UserID,
	)

	if err != nil {

		return nil, exception.Internal(
			"failed to send verification code",
		)
	}

	if user == nil {

		return nil, exception.NotFound(
			"user not found",
		)
	}

	if user.EmailVerified {

		return nil, exception.BadRequest(
			"email already verified",
		)
	}

	// =====================================================
	// COOLDOWN
	// =====================================================

	cooldownKey := cachekey.EmailVerificationCooldown(
		user.ID,
	)

	if u.cache.Exists(
		ctx,
		cooldownKey,
	) {

		return nil, exception.BadRequest(
			"please wait before requesting another verification code",
		)
	}

	// =====================================================
	// GENERATE CODE
	// =====================================================

	n, err := rand.Int(
		rand.Reader,
		big.NewInt(
			1000000,
		),
	)

	if err != nil {

		return nil, exception.Internal(
			"failed to send verification code",
		)
	}

	code := fmt.Sprintf(
		"%06d",
		n.Int64(),
	)

	// =====================================================
	// SAVE CODE
	// =====================================================

	err = u.cache.Set(
		ctx,
		cachekey.EmailVerification(
			user.ID,
		),
		code,
		u.emailVerificationTTL,
	)

	if err != nil {

		return nil, exception.Internal(
			"failed to send verification code",
		)
	}

	// =====================================================
	// SAVE COOLDOWN
	// =====================================================

	err = u.cache.Set(
		ctx,
		cooldownKey,
		"1",
		u.emailVerificationCooldownTTL,
	)

	if err != nil {

		return nil, exception.Internal(
			"failed to send verification code",
		)
	}

	// =====================================================
	// EMAIL
	// =====================================================

	html := template.EmailVerificationCode(
		code,
	)

	err = u.mailer.Send(
		user.Email,
		"Verify Your Email",
		html,
	)

	if err != nil {

		return nil, exception.Internal(
			"failed to send verification code",
		)
	}

	// =====================================================
	// RESPONSE
	// =====================================================

	return &model.WebResponse[any]{
		Message: "verification code sent successfully",
	}, nil
}
