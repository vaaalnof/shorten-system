package report

import (
	"context"

	"shortener-service/internal/delivery/http/middleware"
	"shortener-service/internal/exception"
	"shortener-service/internal/model"
	"shortener-service/internal/model/converter"
)

// =====================================================
// REPORT DEVICES
// =====================================================

func (u *ReportUseCase) Devices(
	ctx context.Context,
	request *model.GetReportDevicesRequest,
) (*model.WebResponse[*model.ReportDevicesResponse], error) {

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
			"failed to retrieve report devices",
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

	items, err := u.reportRepo.GetDevices(
		ctx,
		url.ID,
	)

	if err != nil {

		return nil, exception.Internal(
			"failed to retrieve report devices",
		)
	}

	return &model.WebResponse[*model.ReportDevicesResponse]{
		Message: "report devices retrieved successfully",
		Data: converter.ToReportDevicesResponse(
			items,
		),
	}, nil
}
