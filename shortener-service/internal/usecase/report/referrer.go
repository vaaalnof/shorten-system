package report

import (
	"context"

	"shortener-service/internal/delivery/http/middleware"
	"shortener-service/internal/exception"
	"shortener-service/internal/model"
	"shortener-service/internal/model/converter"
)

// =====================================================
// REPORT REFERRERS
// =====================================================

func (u *ReportUseCase) Referrers(
	ctx context.Context,
	request *model.GetReportReferrersRequest,
) (*model.WebResponse[*model.ReportReferrersResponse], error) {

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
			"failed to retrieve report referrers",
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

	items, err := u.reportRepo.GetReferrers(
		ctx,
		url.ID,
	)

	if err != nil {

		return nil, exception.Internal(
			"failed to retrieve report referrers",
		)
	}

	return &model.WebResponse[*model.ReportReferrersResponse]{
		Message: "report referrers retrieved successfully",
		Data: converter.ToReportReferrersResponse(
			items,
		),
	}, nil
}
