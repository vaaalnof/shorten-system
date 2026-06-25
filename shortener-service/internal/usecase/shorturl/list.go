package shorturl

import (
	"context"
	"math"

	"shortener-service/internal/delivery/http/middleware"
	"shortener-service/internal/exception"
	"shortener-service/internal/model"
	"shortener-service/internal/model/converter"
)

const (
	DefaultPageSize = 20
	MaxPageSize     = 50
)

// =====================================================
// LIST URLS
// =====================================================

func (u *URLUseCase) List(
	ctx context.Context,
	request *model.ListURLsRequest,
) (*model.WebResponse[[]*model.ListURLResponse], error) {

	meta := middleware.GetMeta(
		ctx,
	)

	if meta == nil ||
		meta.UserID == "" {

		return nil, exception.Unauthorized(
			"unauthorized",
		)
	}

	page := request.Page

	if page <= 0 {
		page = 1
	}

	size := request.Size

	if size <= 0 {
		size = DefaultPageSize
	}

	if size > MaxPageSize {
		size = MaxPageSize
	}

	offset := (page - 1) * size

	totalItem, err := u.urlRepo.CountByUserID(
		ctx,
		meta.UserID,
	)

	if err != nil {

		return nil, exception.Internal(
			"failed to retrieve shorturls",
		)
	}

	urls, err := u.urlRepo.ListByUserID(
		ctx,
		meta.UserID,
		size,
		offset,
	)

	if err != nil {

		return nil, exception.Internal(
			"failed to retrieve shorturls",
		)
	}

	totalPage := int64(0)

	if totalItem > 0 {

		totalPage = int64(
			math.Ceil(
				float64(totalItem) /
					float64(size),
			),
		)
	}

	return &model.WebResponse[[]*model.ListURLResponse]{
		Message: "shorturls retrieved successfully",
		Data: converter.ToListURLResponses(
			urls,
		),
		Paging: &model.PageMetadata{
			Page:      page,
			Size:      size,
			TotalItem: totalItem,
			TotalPage: totalPage,
		},
	}, nil
}
