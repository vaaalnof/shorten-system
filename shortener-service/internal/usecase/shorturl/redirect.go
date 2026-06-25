package shorturl

import (
	"context"
	"strings"

	"shortener-service/internal/delivery/http/middleware"
	"shortener-service/internal/exception"
	"shortener-service/internal/model"
	cachekey "shortener-service/internal/utils/cache"
)

func (u *URLUseCase) Redirect(
	ctx context.Context,
	request *model.RedirectURLRequest,
) (string, error) {

	if err := u.validate.Struct(
		request,
	); err != nil {

		return "", exception.Validation(
			err,
		)
	}

	shortCode := strings.ToLower(
		strings.TrimSpace(
			request.ShortCode,
		),
	)

	meta := middleware.GetMeta(
		ctx,
	)

	cacheKey := cachekey.URL(
		shortCode,
	)

	// =====================================================
	// CACHE
	// =====================================================

	if url, found := u.getCachedURL(
		ctx,
		cacheKey,
	); found {

		u.log.WithField(
			"short_code",
			shortCode,
		).Debug(
			"shorturl cache hit",
		)

		if url.PasswordHash != nil {

			return "", exception.Forbidden(
				"password required",
			)
		}

		u.analyticsPublisher.PublishClick(
			meta,
			url,
		)

		return url.OriginalURL, nil
	}

	u.log.WithField(
		"short_code",
		shortCode,
	).Debug(
		"shorturl cache miss",
	)

	// =====================================================
	// DATABASE
	// =====================================================

	url, err := u.urlRepo.FindByShortCode(
		ctx,
		shortCode,
	)

	if err != nil {

		return "", exception.Internal(
			"failed to get shorturl",
		)
	}

	if err := u.validateURL(
		url,
	); err != nil {

		return "", err
	}

	// =====================================================
	// PASSWORD
	// =====================================================

	if url.PasswordHash != nil {

		return "", exception.Forbidden(
			"password required",
		)
	}

	// =====================================================
	// CACHE SET
	// =====================================================

	u.cacheURL(
		ctx,
		cacheKey,
		url,
	)

	// =====================================================
	// ANALYTICS
	// =====================================================

	u.analyticsPublisher.PublishClick(
		meta,
		url,
	)

	return url.OriginalURL, nil
}
