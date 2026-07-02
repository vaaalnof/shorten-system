package report

import (
	"context"

	"shortener-service/internal/delivery/http/middleware"
	"shortener-service/internal/exception"
	"shortener-service/internal/model"
	"shortener-service/internal/model/converter"
)

// =====================================================
// REPORT BROWSERS
// =====================================================

func (u *ReportUseCase) Browsers(
	ctx context.Context,
	request *model.GetReportBrowsersRequest,
) (*model.WebResponse[*model.ReportBrowsersResponse], error) {

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
			"failed to retrieve report browsers",
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

	items, err := u.reportRepo.GetBrowsers(
		ctx,
		url.ID,
	)

	if err != nil {

		return nil, exception.Internal(
			"failed to retrieve report browsers",
		)
	}

	return &model.WebResponse[*model.ReportBrowsersResponse]{
		Message: "report browsers retrieved successfully",
		Data: converter.ToReportBrowsersResponse(
			items,
		),
	}, nil
}
