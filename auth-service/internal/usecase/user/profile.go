package user

import (
	"context"

	"auth-service/internal/exception"
	"auth-service/internal/model"
	"auth-service/internal/model/converter"
	"auth-service/internal/usecase/port"
)

type ProfileUseCase struct {
	*UseCase

	userRepo port.UserRepository
}

func NewProfileUseCase(
	base *UseCase,
	userRepo port.UserRepository,
) *ProfileUseCase {

	return &ProfileUseCase{
		UseCase:  base,
		userRepo: userRepo,
	}
}

func (u *ProfileUseCase) Execute(
	ctx context.Context,
) (*model.WebResponse[*model.ProfileResponse], error) {

	authUser, err := u.GetAuthenticatedUser(
		ctx,
	)

	if err != nil {
		return nil, err
	}

	user, err := u.userRepo.FindByUserID(
		ctx,
		authUser.UserID,
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

	if !user.IsActive {

		return nil, exception.Unauthorized(
			"account is inactive",
		)
	}

	return &model.WebResponse[*model.ProfileResponse]{
		Message: "profile fetched successfully",
		Data: converter.ToProfileResponse(
			user,
		),
	}, nil
}
