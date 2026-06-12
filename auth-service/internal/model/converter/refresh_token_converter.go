package converter

import "auth-service/internal/model"

func ToRefreshTokenResponse(
	accessToken string,
	refreshToken string,
) *model.RefreshTokenResponse {

	return &model.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}
