package entity

type User struct {
	ID              string
	Email           string
	FirstName       string
	LastName        string
	AvatarURL       *string
	IsActive        bool
	EmailVerified   bool
	EmailVerifiedAt *int64
	CreatedAt       int64
	UpdatedAt       *int64
}
