package shorturl

import (
	"context"
	"encoding/json"
	"shortener-service/internal/entity"
	"shortener-service/internal/model"
)

func (u *URLUseCase) getCachedURL(
	ctx context.Context,
	cacheKey string,
) (*entity.URL, bool) {

	cachedValue, err := u.cache.Get(
		ctx,
		cacheKey,
	)

	if err != nil ||
		cachedValue == "" {

		return nil, false
	}

	var cached model.CachedURL

	if err := json.Unmarshal(
		[]byte(cachedValue),
		&cached,
	); err != nil {

		u.log.WithError(
			err,
		).Warn(
			"failed to unmarshal cached shorturl",
		)

		return nil, false
	}

	url := &entity.URL{
		ID:           cached.ID,
		ShortCode:    cached.ShortCode,
		OriginalURL:  cached.OriginalURL,
		IsActive:     cached.IsActive,
		PasswordHash: cached.PasswordHash,
		ExpiredAt:    cached.ExpiredAt,
	}

	return url, true
}

func (u *URLUseCase) cacheURL(
	ctx context.Context,
	cacheKey string,
	url *entity.URL,
) {

	if url == nil {
		return
	}

	cached := model.CachedURL{
		ID:           url.ID,
		ShortCode:    url.ShortCode,
		OriginalURL:  url.OriginalURL,
		IsActive:     url.IsActive,
		PasswordHash: url.PasswordHash,
		ExpiredAt:    url.ExpiredAt,
	}

	payload, err := json.Marshal(
		cached,
	)

	if err != nil {
		return
	}

	if err := u.cache.Set(
		ctx,
		cacheKey,
		string(payload),
		u.urlCacheTTL,
	); err != nil {

		u.log.WithError(
			err,
		).Warn(
			"failed to cache shorturl",
		)
	}
}

func (u *URLUseCase) deleteCachedURL(
	ctx context.Context,
	cacheKey string,
) {

	if err := u.cache.Delete(
		ctx,
		cacheKey,
	); err != nil {

		u.log.WithError(
			err,
		).Warn(
			"failed to delete shorturl cache",
		)
	}
}
