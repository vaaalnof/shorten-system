package entity

type URLDailyAnalytics struct {
	ID             int64
	UrlID          string
	AnalyticsDate  int64
	TotalClicks    int64
	UniqueVisitors int64
	CreatedAt      int64
	UpdatedAt      int64
}
