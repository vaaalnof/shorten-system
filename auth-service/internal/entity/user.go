package entity

type User struct {
	ID            string
	Email         string
	FirstName     string
	LastName      string
	AvatarURL     *string
	IsActive      bool
	EmailVerified bool
	CreatedAt     int64
	UpdatedAt     *int64
}
