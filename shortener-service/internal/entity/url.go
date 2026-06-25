package entity

type URL struct {
	ID           string
	UserID       string
	ShortCode    string
	OriginalURL  string
	IsActive     bool
	PasswordHash *string
	ExpiredAt    *int64
	DeletedAt    *int64
	CreatedAt    int64
	UpdatedAt    int64
}
