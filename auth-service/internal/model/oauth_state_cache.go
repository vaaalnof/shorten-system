package model

type OAuthStateCache struct {
	Provider  string `json:"provider"`
	CreatedAt int64  `json:"created_at"`
}
