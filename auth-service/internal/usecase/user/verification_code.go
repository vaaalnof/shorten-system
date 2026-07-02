package user

import (
	"context"

	"auth-service/internal/exception"
	"auth-service/internal/model"
	"auth-service/internal/model/converter"
	"auth-service/internal/usecase/port"
	cachekey "auth-service/internal/utils/cache"
)

type VerificationCodeUseCase struct {
	*UseCase

	userRepo port.UserRepository
	cache    port.Cache
}

func NewVerificationCodeUseCase(
	base *UseCase,
	userRepo port.UserRepository,
	cache port.Cache,
) *VerificationCodeUseCase {

	return &VerificationCodeUseCase{
		UseCase:  base,
		userRepo: userRepo,
		cache:    cache,
	}
}

func (u *VerificationCodeUseCase) Execute(
	ctx context.Context,
	request *model.VerificationCodeRequest,
) (*model.WebResponse[*model.VerificationCodeResponse], error) {

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
			"failed to verify email",
		)
	}

	if user == nil {

		return nil, exception.NotFound(
			"user not found",
		)
	}

	// =====================================================
	// ALREADY VERIFIED
	// =====================================================

	if user.EmailVerified {

		return nil, exception.BadRequest(
			"email already verified",
		)
	}

	// =====================================================
	// REDIS CODE
	// =====================================================

	cacheCode, err := u.cache.Get(
		ctx,
		cachekey.EmailVerification(
			user.ID,
		),
	)

	if err != nil {

		return nil, exception.Internal(
			"failed to verify email",
		)
	}

	if cacheCode == "" {

		return nil, exception.BadRequest(
			"verification code expired",
		)
	}

	// =====================================================
	// VALIDATE CODE
	// =====================================================

	if cacheCode != request.Code {

		return nil, exception.BadRequest(
			"invalid verification code",
		)
	}

	// =====================================================
	// VERIFY EMAIL
	// =====================================================

	emailVerifiedAt, err := u.userRepo.UpdateEmailVerified(
		ctx,
		user.ID,
	)

	if err != nil {

		return nil, exception.Internal(
			"failed to verify email",
		)
	}

	user.EmailVerified = true
	user.EmailVerifiedAt = &emailVerifiedAt

	// =====================================================
	// DELETE REDIS CODE
	// =====================================================

	_ = u.cache.Delete(
		ctx,
		cachekey.EmailVerification(
			user.ID,
		),
	)

	// =====================================================
	// RESPONSE
	// =====================================================

	return &model.WebResponse[*model.VerificationCodeResponse]{
		Message: "email verified successfully",
		Data: converter.ToVerificationCodeResponse(
			user,
		),
	}, nil
}
