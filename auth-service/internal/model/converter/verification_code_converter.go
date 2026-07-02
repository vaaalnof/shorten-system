package converter

import (
	"auth-service/internal/entity"
	"auth-service/internal/model"
	"auth-service/internal/utils"
)

func ToVerificationCodeResponse(
	user *entity.User,
) *model.VerificationCodeResponse {

	var verifiedAt string

	if user.EmailVerifiedAt != nil {

		verifiedAt = utils.FormatUnixTime(
			*user.EmailVerifiedAt,
			"2006-01-02 15:04:05",
		)
	}

	return &model.VerificationCodeResponse{
		ID:              user.ID,
		EmailVerifiedAt: verifiedAt,
	}
}
