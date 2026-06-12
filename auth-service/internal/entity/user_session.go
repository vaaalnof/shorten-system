package entity

type UserSession struct {
	ID           string
	UserID       string
	RefreshToken string
	IPAddress    *string
	UserAgent    *string
	ExpiredAt    int64
	RevokedAt    *int64
	CreatedAt    int64
}
