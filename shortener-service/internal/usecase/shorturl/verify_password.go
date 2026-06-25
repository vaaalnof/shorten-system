package shorturl

import (
	"context"
	"strings"

	"shortener-service/internal/delivery/http/middleware"
	"shortener-service/internal/exception"
	"shortener-service/internal/model"
)

// =====================================================
// VERIFY PASSWORD
// =====================================================

func (u *URLUseCase) VerifyPassword(
	ctx context.Context,
	request *model.VerifyURLPasswordRequest,
) (string, error) {

	if err := u.validate.Struct(
		request,
	); err != nil {

		return "", exception.Validation(
			err,
		)
	}

	shortCode, err := u.validateShortCode(
		request.ShortCode,
	)

	if err != nil {

		return "", err
	}

	password := strings.TrimSpace(
		request.Password,
	)

	if password == "" {

		return "", exception.BadRequest(
			"password is required",
		)
	}

	// =====================================================
	// DATABASE
	// =====================================================

	url, err := u.urlRepo.FindByShortCode(
		ctx,
		shortCode,
	)

	if err != nil {

		return "", exception.Internal(
			"failed to verify shorturl password",
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

	if url.PasswordHash == nil {

		return "", exception.BadRequest(
			"shorturl does not have a password",
		)
	}

	if err := u.passwordHash.Compare(
		*url.PasswordHash,
		password,
	); err != nil {

		return "", exception.BadRequest(
			"invalid password",
		)
	}

	// =====================================================
	// ANALYTICS
	// =====================================================

	meta := middleware.GetMeta(
		ctx,
	)

	u.analyticsPublisher.PublishClick(
		meta,
		url,
	)

	return url.OriginalURL, nil
}
