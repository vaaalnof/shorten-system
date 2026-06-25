package shorturl

import (
	"strings"
	"time"

	"shortener-service/internal/entity"
	"shortener-service/internal/exception"
)

func (u *URLUseCase) validateShortCode(
	shortCode string,
) (string, error) {

	shortCode = strings.TrimSpace(
		shortCode,
	)

	if shortCode == "" {

		return "", exception.BadRequest(
			"short code is required",
		)
	}

	return shortCode, nil
}

func (u *URLUseCase) validateURL(
	url *entity.URL,
) error {

	if url == nil {

		return exception.NotFound(
			"shorturl not found",
		)
	}

	if !url.IsActive {

		return exception.NotFound(
			"shorturl not found",
		)
	}

	if url.ExpiredAt != nil {

		now := time.Now().Unix()

		if *url.ExpiredAt <= now {

			return exception.NotFound(
				"shorturl expired",
			)
		}
	}

	return nil
}
