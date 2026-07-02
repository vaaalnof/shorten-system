package converter

import "auth-service/internal/model"

func ToGoogleLoginResponse(
	loginURL string,
) *model.GoogleLoginResponse {

	return &model.GoogleLoginResponse{
		LoginURL: loginURL,
	}
}
