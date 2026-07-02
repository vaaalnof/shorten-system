package report

import (
	"context"

	"shortener-service/internal/delivery/http/middleware"
	"shortener-service/internal/exception"
	"shortener-service/internal/model"
	"shortener-service/internal/model/converter"
)

// =====================================================
// TOP LINKS
// =====================================================

func (u *ReportUseCase) TopLinks(
	ctx context.Context,
	request *model.GetTopLinksRequest,
) (*model.WebResponse[*model.TopLinksResponse], error) {

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

	items, err := u.reportRepo.GetTopLinks(
		ctx,
		meta.UserID,
	)

	if err != nil {

		return nil, exception.Internal(
			"failed to retrieve top links",
		)
	}

	return &model.WebResponse[*model.TopLinksResponse]{
		Message: "top links retrieved successfully",
		Data: converter.ToTopLinksResponse(
			items,
		),
	}, nil
}
