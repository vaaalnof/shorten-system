package shorturl

import (
	"context"

	"shortener-service/internal/delivery/http/middleware"
	"shortener-service/internal/exception"
	"shortener-service/internal/model"
	"shortener-service/internal/model/converter"
	cachekey "shortener-service/internal/utils/cache"
)

// =====================================================
// DELETE URL
// =====================================================

func (u *URLUseCase) Delete(
	ctx context.Context,
	request *model.DeleteURLRequest,
) (*model.WebResponse[*model.DeleteURLResponse], error) {

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
			"failed to delete shorturl",
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
	// DELETE
	// =====================================================

	if err := u.urlRepo.Delete(
		ctx,
		url,
	); err != nil {

		return nil, exception.Internal(
			"failed to delete shorturl",
		)
	}

	// =====================================================
	// CACHE INVALIDATION
	// =====================================================

	u.deleteCachedURL(
		ctx,
		cachekey.URL(
			url.ShortCode,
		),
	)

	return &model.WebResponse[*model.DeleteURLResponse]{
		Message: "shorturl deleted successfully",
		Data: converter.ToDeleteURLResponse(
			url,
		),
	}, nil
}
