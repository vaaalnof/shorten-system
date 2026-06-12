package converter

import "auth-service/internal/model"

func ToLoginResponse(
	accessToken string,
	refreshToken string,
) *model.LoginResponse {
	return &model.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}
