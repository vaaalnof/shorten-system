package report

import (
	"context"

	"shortener-service/internal/delivery/http/middleware"
	"shortener-service/internal/exception"
	"shortener-service/internal/model"
	"shortener-service/internal/model/converter"
)

// =====================================================
// REPORT COUNTRIES
// =====================================================

func (u *ReportUseCase) Countries(
	ctx context.Context,
	request *model.GetReportCountriesRequest,
) (*model.WebResponse[*model.ReportCountriesResponse], error) {

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
			"failed to retrieve report countries",
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

	items, err := u.reportRepo.GetCountries(
		ctx,
		url.ID,
	)

	if err != nil {

		return nil, exception.Internal(
			"failed to retrieve report countries",
		)
	}

	return &model.WebResponse[*model.ReportCountriesResponse]{
		Message: "report countries retrieved successfully",
		Data: converter.ToReportCountriesResponse(
			items,
		),
	}, nil
}
