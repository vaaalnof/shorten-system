package model

type ValidateTokenResponse struct {
	UserID    string `json:"user_id"`
	SessionID string `json:"session_id"`
	IsActive  bool   `json:"is_active"`
}
