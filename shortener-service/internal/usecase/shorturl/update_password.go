package shorturl

import (
	"context"
	cachekey "shortener-service/internal/utils/cache"
	"strings"

	"shortener-service/internal/delivery/http/middleware"
	"shortener-service/internal/exception"
	"shortener-service/internal/model"
	"shortener-service/internal/model/converter"
)

// =====================================================
// UPDATE PASSWORD
// =====================================================

func (u *URLUseCase) UpdatePassword(
	ctx context.Context,
	request *model.UpdateURLPasswordRequest,
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
			"failed to update shorturl password",
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

	oldPassword := strings.TrimSpace(
		request.OldPassword,
	)

	newPassword := strings.TrimSpace(
		request.NewPassword,
	)

	// =====================================================
	// NEW PASSWORD
	// =====================================================

	if len(newPassword) < 7 {

		return nil, exception.BadRequest(
			"new password minimum length is 7",
		)
	}

	// =====================================================
	// EXISTING PASSWORD
	// =====================================================

	if url.PasswordHash != nil {

		if oldPassword == "" {

			return nil, exception.BadRequest(
				"old password is required",
			)
		}

		if len(oldPassword) < 7 {

			return nil, exception.BadRequest(
				"old password minimum length is 7",
			)
		}

		if err := u.passwordHash.Compare(
			*url.PasswordHash,
			oldPassword,
		); err != nil {

			return nil, exception.BadRequest(
				"old password is invalid",
			)
		}

		if oldPassword == newPassword {

			return nil, exception.BadRequest(
				"new password must be different from old password",
			)
		}
	}

	// =====================================================
	// HASH PASSWORD
	// =====================================================

	passwordHash, err := u.passwordHash.Hash(
		newPassword,
	)

	if err != nil {

		return nil, exception.Internal(
			"failed to hash password",
		)
	}

	url.PasswordHash = &passwordHash

	if err := u.urlRepo.UpdatePassword(
		ctx,
		url,
	); err != nil {

		return nil, exception.Internal(
			"failed to update shorturl password",
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
		Message: "shorturl password updated successfully",
		Data: converter.ToUpdateURLPasswordResponse(
			url,
		),
	}, nil
}
