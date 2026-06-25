package shorturl

import (
	"context"
	"strings"

	"shortener-service/internal/delivery/http/middleware"
	"shortener-service/internal/exception"
	"shortener-service/internal/model"
	"shortener-service/internal/model/converter"
	cachekey "shortener-service/internal/utils/cache"
)

// =====================================================
// REMOVE PASSWORD
// =====================================================

func (u *URLUseCase) RemovePassword(
	ctx context.Context,
	request *model.RemoveURLPasswordRequest,
) (*model.WebResponse[*model.UpdateURLPasswordResponse], error) {

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
			"failed to remove shorturl password",
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

	if url.PasswordHash == nil {

		return nil, exception.BadRequest(
			"shorturl does not have a password",
		)
	}

	oldPassword := strings.TrimSpace(
		request.OldPassword,
	)

	if err := u.passwordHash.Compare(
		*url.PasswordHash,
		oldPassword,
	); err != nil {

		return nil, exception.BadRequest(
			"old password is invalid",
		)
	}

	if err := u.urlRepo.RemovePassword(
		ctx,
		url,
	); err != nil {

		return nil, exception.Internal(
			"failed to remove shorturl password",
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

	return &model.WebResponse[*model.UpdateURLPasswordResponse]{
		Message: "shorturl password removed successfully",
		Data: converter.ToUpdateURLPasswordResponse(
			url,
		),
	}, nil
}
