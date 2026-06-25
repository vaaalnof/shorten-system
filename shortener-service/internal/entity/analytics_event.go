package entity

type AnalyticsEvent struct {
	ID        int64
	UrlID     string
	ShortCode string
	Referer   *string
	Source    *string
	IPAddress *string
	UserAgent *string
	Browser   *string
	OS        *string
	Device    *string
	Country   *string
	ClickedAt int64
}
