package converter

import "auth-service/internal/model"

func ToValidateTokenResponse(
	userID string,
	sessionID string,
	emailVerified bool,
) *model.ValidateTokenResponse {

	status := "not_verified"

	if emailVerified {
		status = "verified"
	}

	return &model.ValidateTokenResponse{
		UserID:        userID,
		SessionID:     sessionID,
		EmailVerified: status,
	}
}
