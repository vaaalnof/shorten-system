package model

type ValidateTokenResponse struct {
	UserID        string `json:"user_id"`
	SessionID     string `json:"session_id"`
	EmailVerified string `json:"email_verified"`
}
