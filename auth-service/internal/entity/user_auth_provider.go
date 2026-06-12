package entity

type AuthProvider string

const (
	AuthProviderLocal  AuthProvider = "local"
	AuthProviderGoogle AuthProvider = "google"
	AuthProviderGithub AuthProvider = "github"
)

type UserAuthProvider struct {
	ID             string
	UserID         string
	Provider       AuthProvider
	ProviderUserID *string
	PasswordHash   *string
	CreatedAt      int64
}
