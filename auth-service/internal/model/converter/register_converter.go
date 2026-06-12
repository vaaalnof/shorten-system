package converter

import (
	"auth-service/internal/entity"
	"auth-service/internal/model"
	"auth-service/internal/utils"
)

func ToRegisterUserResponse(
	u *entity.User,
) *model.RegisterUserResponse {

	return &model.RegisterUserResponse{
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  &u.LastName,
		CreatedAt: utils.FormatUnixTime(u.CreatedAt, "2006-01-02 15:04:05"),
	}
}
