package model

type ProfileResponse struct {
	ID            string  `json:"id"`
	Email         string  `json:"email"`
	FirstName     string  `json:"first_name"`
	LastName      *string `json:"last_name"`
	AvatarURL     *string `json:"avatar_url"`
	IsActive      bool    `json:"is_active"`
	EmailVerified bool    `json:"email_verified"`
	CreatedAt     string  `json:"created_at"`
}
