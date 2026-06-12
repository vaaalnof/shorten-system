package converter

import "auth-service/internal/model"

func ToValidateTokenResponse(
	userID string,
	sessionID string,
) *model.ValidateTokenResponse {

	return &model.ValidateTokenResponse{
		UserID:    userID,
		SessionID: sessionID,
	}
}
