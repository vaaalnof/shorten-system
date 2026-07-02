package converter

import (
	"auth-service/internal/entity"
	"auth-service/internal/model"
	"auth-service/internal/utils"
)

func ToProfileResponse(
	u *entity.User,
) *model.ProfileResponse {

	var lastName *string

	if u.LastName != "" {
		lastName = &u.LastName
	}

	return &model.ProfileResponse{
		ID:            u.ID,
		Email:         u.Email,
		FirstName:     u.FirstName,
		LastName:      lastName,
		AvatarURL:     u.AvatarURL,
		IsActive:      u.IsActive,
		EmailVerified: u.EmailVerified,
		CreatedAt: utils.FormatUnixTime(
			u.CreatedAt,
			"2006-01-02 15:04:05",
		),
	}
}
