package auth

import (
	"context"

	"auth-service/internal/model"
	"auth-service/internal/model/converter"
)

type ValidateTokenUseCase struct {
	*AccessSessionUseCase
}

func NewValidateTokenUseCase(
	base *AccessSessionUseCase,
) *ValidateTokenUseCase {

	return &ValidateTokenUseCase{
		AccessSessionUseCase: base,
	}
}

func (u *ValidateTokenUseCase) Execute(
	ctx context.Context,
) (*model.WebResponse[*model.ValidateTokenResponse], error) {

	// =====================================================
	// ACCESS SESSION
	// =====================================================

	session, err := u.GetSession(
		ctx,
	)

	if err != nil {
		return nil, err
	}

	// =====================================================
	// RESPONSE
	// =====================================================

	return &model.WebResponse[*model.ValidateTokenResponse]{
		Message: "token valid",
		Data: converter.ToValidateTokenResponse(
			session.UserID,
			session.SessionID,
			session.EmailVerified,
		),
	}, nil
}
