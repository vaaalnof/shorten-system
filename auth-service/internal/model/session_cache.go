package model

type SessionCache struct {
	UserID    string `json:"user_id"`
	SessionID string `json:"session_id"`
	ExpiredAt int64  `json:"expired_at"`
}
