package shorturl

import (
	"context"
	"strings"
	"time"

	"shortener-service/internal/delivery/http/middleware"
	"shortener-service/internal/exception"
	"shortener-service/internal/model"
	"shortener-service/internal/model/converter"
	cachekey "shortener-service/internal/utils/cache"
)

// =====================================================
// UPDATE URL
// =====================================================

func (u *URLUseCase) Update(
	ctx context.Context,
	request *model.UpdateURLRequest,
) (*model.WebResponse[*model.UpdateURLResponse], error) {

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
			"failed to update shorturl",
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

	oldShortCode := url.ShortCode

	// =====================================================
	// ORIGINAL URL
	// =====================================================

	if request.OriginalURL != nil {

		url.OriginalURL = strings.TrimSpace(
			*request.OriginalURL,
		)
	}

	// =====================================================
	// SHORT CODE
	// =====================================================

	if request.ShortCode != nil {

		shortCode := strings.ToLower(
			strings.TrimSpace(
				*request.ShortCode,
			),
		)

		if shortCode != "" &&
			shortCode != url.ShortCode {

			reserved, err := u.reservedAliasRepo.Exists(
				ctx,
				shortCode,
			)

			if err != nil {

				return nil, exception.Internal(
					"failed to update shorturl",
				)
			}

			if reserved {

				return nil, exception.Conflict(
					"short code is reserved",
				)
			}

			existing, err := u.urlRepo.FindByShortCode(
				ctx,
				shortCode,
			)

			if err != nil {

				return nil, exception.Internal(
					"failed to update shorturl",
				)
			}

			if existing != nil &&
				existing.ID != url.ID {

				return nil, exception.Conflict(
					"short code already exists",
				)
			}

			url.ShortCode = shortCode
		}
	}

	// =====================================================
	// STATUS
	// =====================================================

	if request.IsActive != nil {

		url.IsActive = *request.IsActive
	}

	// =====================================================
	// EXPIRED AT
	// =====================================================

	if request.ExpiredAt != nil {

		if *request.ExpiredAt < 0 {

			return nil, exception.BadRequest(
				"expired_at must be greater than or equal to 0",
			)
		}

		if *request.ExpiredAt == 0 {

			url.ExpiredAt = nil

		} else {

			now := time.Now().Unix()

			if *request.ExpiredAt <= now {

				return nil, exception.BadRequest(
					"expired_at must be greater than current time",
				)
			}

			url.ExpiredAt = request.ExpiredAt
		}
	}

	// =====================================================
	// UPDATE
	// =====================================================

	if err := u.urlRepo.Update(
		ctx,
		url,
	); err != nil {

		u.log.WithError(err).Error(
			"failed to update shorturl",
		)

		return nil, exception.Internal(
			"failed to update shorturl",
		)
	}

	// =====================================================
	// CACHE INVALIDATION
	// =====================================================

	u.deleteCachedURL(
		ctx,
		cachekey.URL(
			oldShortCode,
		),
	)

	if oldShortCode != url.ShortCode {

		u.deleteCachedURL(
			ctx,
			cachekey.URL(
				url.ShortCode,
			),
		)
	}

	return &model.WebResponse[*model.UpdateURLResponse]{
		Message: "shorturl updated successfully",
		Data: converter.ToUpdateURLResponse(
			url,
		),
	}, nil
}
