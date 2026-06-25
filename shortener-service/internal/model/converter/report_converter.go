package converter

import (
	"shortener-service/internal/entity"
	"shortener-service/internal/model"
	"shortener-service/internal/utils"
)

func ToReportSummaryResponse(
	summary *entity.ReportSummary,
) *model.ReportSummaryResponse {

	var lastClickAt *string

	if summary.LastClickAt != nil {

		formatted := utils.FormatUnixTime(
			*summary.LastClickAt,
			"2006-01-02 15:04:05",
		)

		lastClickAt = &formatted
	}

	return &model.ReportSummaryResponse{
		TotalClicks:    summary.TotalClicks,
		UniqueVisitors: summary.UniqueVisitors,
		TodayClicks:    summary.TodayClicks,
		TodayVisitors:  summary.TodayVisitors,
		LastClickAt:    lastClickAt,
	}
}

func ToReportChartResponse(

	items []*entity.ReportChart,

) *model.ReportChartResponse {

	result := make(
		[]*model.ReportChartItemResponse,
		0,
		len(items),
	)
	for _, item := range items {
		result = append(
			result,
			&model.ReportChartItemResponse{
				Date: utils.FormatUnixTime(
					item.AnalyticsDate,
					"2006-01-02",
				),
				Clicks:         item.TotalClicks,
				UniqueVisitors: item.UniqueVisitors,
			},
		)
	}
	return &model.ReportChartResponse{
		Items: result,
	}

}

func ToReportReferrersResponse(
	items []*entity.ReportReferrer,
) *model.ReportReferrersResponse {

	result := make(
		[]*model.ReportReferrerItemResponse,
		0,
		len(items),
	)

	for _, item := range items {

		result = append(
			result,
			&model.ReportReferrerItemResponse{
				Referrer: item.Referrer,
				Clicks:   item.Clicks,
			},
		)
	}

	return &model.ReportReferrersResponse{
		Items: result,
	}
}

func ToReportCountriesResponse(
	items []*entity.ReportCountry,
) *model.ReportCountriesResponse {

	result := make(
		[]*model.ReportCountryItemResponse,
		0,
		len(items),
	)

	for _, item := range items {

		result = append(
			result,
			&model.ReportCountryItemResponse{
				Country: item.Country,
				Clicks:  item.Clicks,
			},
		)
	}

	return &model.ReportCountriesResponse{
		Items: result,
	}
}

// =====================================================
// REPORT DEVICES
// =====================================================

func ToReportDevicesResponse(
	items []*entity.ReportDevice,
) *model.ReportDevicesResponse {

	result := make(
		[]*model.ReportDeviceItemResponse,
		0,
		len(items),
	)

	for _, item := range items {

		result = append(
			result,
			&model.ReportDeviceItemResponse{
				Device: item.Device,
				Clicks: item.Clicks,
			},
		)
	}

	return &model.ReportDevicesResponse{
		Items: result,
	}
}
