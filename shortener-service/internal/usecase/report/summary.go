package report

import (
	"context"

	"shortener-service/internal/delivery/http/middleware"
	"shortener-service/internal/exception"
	"shortener-service/internal/model"
	"shortener-service/internal/model/converter"
)

func (u *ReportUseCase) Summary(
	ctx context.Context,
	request *model.GetReportSummaryRequest,
) (*model.WebResponse[*model.ReportSummaryResponse], error) {

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

	url, err := u.urlRepo.FindByID(
		ctx,
		request.ID,
	)

	if err != nil {

		return nil, exception.Internal(
			"failed to get report summary",
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

	summary, err := u.reportRepo.GetSummary(
		ctx,
		url.ID,
	)

	if err != nil {

		return nil, exception.Internal(
			"failed to get report summary",
		)
	}

	return &model.WebResponse[*model.ReportSummaryResponse]{
		Message: "report summary retrieved successfully",
		Data: converter.ToReportSummaryResponse(
			summary,
		),
	}, nil
}
