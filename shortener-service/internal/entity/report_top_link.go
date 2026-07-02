package entity

type TopLink struct {
	ID             string
	ShortCode      string
	OriginalURL    string
	Clicks         int64
	UniqueVisitors int64
	CreatedAt      int64
}
