package report

import (
	"context"

	"shortener-service/internal/delivery/http/middleware"
	"shortener-service/internal/exception"
	"shortener-service/internal/model"
	"shortener-service/internal/model/converter"
)

// =====================================================
// REPORT CHART
// =====================================================

func (u *ReportUseCase) Chart(
	ctx context.Context,
	request *model.GetReportChartRequest,
) (*model.WebResponse[*model.ReportChartResponse], error) {

	if err := u.validate.Struct(
		request,
	); err != nil {

		return nil, exception.Validation(
			err,
		)
	}

	meta := middleware.GetMeta(
		ctx,
	)

	if meta == nil ||
		meta.UserID == "" {

		return nil, exception.Unauthorized(
			"unauthorized",
		)
	}

	// =====================================================
	// URL
	// =====================================================

	url, err := u.urlRepo.FindByID(
		ctx,
		request.ID,
	)

	if err != nil {

		return nil, exception.Internal(
			"failed to retrieve report chart",
		)
	}

	if url == nil {

		return nil, exception.NotFound(
			"shorturl not found",
		)
	}

	if url.UserID != meta.UserID {

		return nil, exception.NotFound(
			"shorturl not found",
		)
	}

	// =====================================================
	// REPORT
	// =====================================================

	items, err := u.reportRepo.GetChart(
		ctx,
		url.ID,
	)

	if err != nil {

		return nil, exception.Internal(
			"failed to retrieve report chart",
		)
	}

	return &model.WebResponse[*model.ReportChartResponse]{
		Message: "report chart retrieved successfully",
		Data: converter.ToReportChartResponse(
			items,
		),
	}, nil
}
