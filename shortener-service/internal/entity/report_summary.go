package entity

type ReportSummary struct {
	TotalClicks    int64
	UniqueVisitors int64
	TodayClicks    int64
	TodayVisitors  int64
	LastClickAt    *int64
}
