package model

type SessionCache struct {
	UserID        string `json:"user_id"`
	SessionID     string `json:"session_id"`
	EmailVerified bool   `json:"email_verified"`
	ExpiredAt     int64  `json:"expired_at"`
	RefreshToken  string `json:"refresh_token,omitempty"`
}
