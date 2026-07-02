package model

// =====================================================
// REPORT SUMMARY
// =====================================================

type GetReportSummaryRequest struct {
	ID string `params:"id" validate:"required,uuid"`
}

type ReportSummaryResponse struct {
	TotalClicks int64 `json:"total_clicks"`

	UniqueVisitors int64 `json:"unique_visitors"`

	TodayClicks int64 `json:"today_clicks"`

	TodayVisitors int64 `json:"today_visitors"`

	LastClickAt *string `json:"last_click_at,omitempty"`
}

// =====================================================

// REPORT CHART

// =====================================================

type GetReportChartRequest struct {
	ID string `params:"id" validate:"required,uuid"`
}

type ReportChartItemResponse struct {
	Date           string `json:"date"`
	Clicks         int64  `json:"clicks"`
	UniqueVisitors int64  `json:"unique_visitors"`
}

type ReportChartResponse struct {
	Items []*ReportChartItemResponse `json:"items"`
}

// =====================================================
// REPORT REFERRERS
// =====================================================

type GetReportReferrersRequest struct {
	ID string `params:"id" validate:"required,uuid"`
}

type ReportReferrerItemResponse struct {
	Referrer string `json:"referrer"`
	Clicks   int64  `json:"clicks"`
}

type ReportReferrersResponse struct {
	Items []*ReportReferrerItemResponse `json:"items"`
}

// =====================================================
// REPORT COUNTRIES
// =====================================================

type GetReportCountriesRequest struct {
	ID string `params:"id" validate:"required,uuid"`
}

type ReportCountryItemResponse struct {
	Country string `json:"country"`
	Clicks  int64  `json:"clicks"`
}

type ReportCountriesResponse struct {
	Items []*ReportCountryItemResponse `json:"items"`
}

// =====================================================
// REPORT DEVICES
// =====================================================

type GetReportDevicesRequest struct {
	ID string `params:"id" validate:"required,uuid"`
}

type ReportDeviceItemResponse struct {
	Device string `json:"device"`
	Clicks int64  `json:"clicks"`
}

type ReportDevicesResponse struct {
	Items []*ReportDeviceItemResponse `json:"items"`
}

// =====================================================
// REPORT BROWSERS
// =====================================================

type GetReportBrowsersRequest struct {
	ID string `params:"id" validate:"required,uuid"`
}

type ReportBrowserItemResponse struct {
	Browser string `json:"browser"`
	Clicks  int64  `json:"clicks"`
}

type ReportBrowsersResponse struct {
	Items []*ReportBrowserItemResponse `json:"items"`
}

// =====================================================
// TOP LINKS
// =====================================================

type GetTopLinksRequest struct {
}

type TopLinkItemResponse struct {
	ID             string `json:"id"`
	ShortCode      string `json:"short_code"`
	OriginalURL    string `json:"original_url"`
	Clicks         int64  `json:"clicks"`
	UniqueVisitors int64  `json:"unique_visitors"`
	CreatedAt      string `json:"created_at"`
}

type TopLinksResponse struct {
	Items []*TopLinkItemResponse `json:"items"`
}
