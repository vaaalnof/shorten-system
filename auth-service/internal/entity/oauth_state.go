package entity

type OAuthState struct {
	ID        string
	State     string
	Provider  string
	ExpiredAt int64
	CreatedAt int64
}
