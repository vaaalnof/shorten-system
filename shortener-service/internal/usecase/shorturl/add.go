package shorturl

import (
	"context"
	"strings"
	"time"

	"github.com/google/uuid"

	"shortener-service/internal/delivery/http/middleware"
	"shortener-service/internal/entity"
	"shortener-service/internal/exception"
	"shortener-service/internal/model"
	"shortener-service/internal/model/converter"
)

// =====================================================
// CREATE URL
// =====================================================

func (u *URLUseCase) Add(
	ctx context.Context,
	request *model.AddURLRequest,
) (*model.WebResponse[*model.AddURLResponse], error) {

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

	if meta.EmailVerified != "verified" {

		return nil, exception.Forbidden(
			"email is not verified",
		)
	}

	shortCode := strings.TrimSpace(
		request.ShortCode,
	)

	shortCode = strings.ToLower(
		shortCode,
	)

	if shortCode != "" &&
		len(shortCode) < 3 {

		return nil, exception.BadRequest(
			"short code minimum length is 3",
		)
	}

	if shortCode == "" {

		shortCode = uuid.NewString()[:8]
	}

	password := strings.TrimSpace(
		request.Password,
	)

	if password != "" &&
		len(password) < 7 {

		return nil, exception.BadRequest(
			"password minimum length is 7",
		)
	}

	var passwordHash *string

	if password != "" {

		hash, err := u.passwordHash.Hash(
			password,
		)

		if err != nil {

			return nil, exception.Internal(
				"failed to hash password",
			)
		}

		passwordHash = &hash
	}

	var expiredAt *int64

	if request.ExpiredAt < 0 {

		return nil, exception.BadRequest(
			"expired_at must be greater than or equal to 0",
		)
	}

	if request.ExpiredAt > 0 {

		now := time.Now().Unix()

		if request.ExpiredAt <= now {

			return nil, exception.BadRequest(
				"expired_at must be greater than current time",
			)
		}

		expiredAt = &request.ExpiredAt
	}

	reserved, err := u.reservedAliasRepo.Exists(
		ctx,
		shortCode,
	)

	if err != nil {

		return nil, exception.Internal(
			"failed to create shorturl",
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
			"failed to create shorturl",
		)
	}

	if existing != nil {

		return nil, exception.Conflict(
			"short code already exists",
		)
	}

	url := &entity.URL{
		ID:           uuid.NewString(),
		UserID:       meta.UserID,
		ShortCode:    shortCode,
		OriginalURL:  request.OriginalURL,
		IsActive:     true,
		PasswordHash: passwordHash,
		ExpiredAt:    expiredAt,
	}

	if err := u.urlRepo.Add(
		ctx,
		url,
	); err != nil {

		return nil, exception.Internal(
			"failed to create shorturl",
		)
	}

	return &model.WebResponse[*model.AddURLResponse]{
		Message: "shorturl created successfully",
		Data: converter.ToAddURLResponse(
			url,
		),
	}, nil
}
