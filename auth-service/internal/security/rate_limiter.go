package security

import (
	"auth-service/internal/exception"
	"auth-service/internal/usecase/port"
	"context"
	"strconv"
	"time"
)

type RateLimiter interface {
	Check(
		ctx context.Context,
		key string,
		maxAttempts int,
	) error

	Increment(
		ctx context.Context,
		key string,
		ttl time.Duration,
	) error

	Reset(
		ctx context.Context,
		key string,
	) error
}

type rateLimiter struct {
	cache port.Cache
}

func NewRateLimiter(
	cache port.Cache,
) RateLimiter {

	return &rateLimiter{
		cache: cache,
	}
}

func (r *rateLimiter) Check(
	ctx context.Context,
	key string,
	maxAttempts int,
) error {

	value, err := r.cache.Get(
		ctx,
		key,
	)

	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	attempts, err := strconv.Atoi(
		value,
	)

	if err != nil {
		return nil
	}

	if attempts >= maxAttempts {

		return exception.TooManyRequests(
			"too many requests, please try again later",
		)
	}

	return nil
}

func (r *rateLimiter) Increment(
	ctx context.Context,
	key string,
	ttl time.Duration,
) error {

	_, err := r.cache.Incr(
		ctx,
		key,
		ttl,
	)

	return err
}

func (r *rateLimiter) Reset(
	ctx context.Context,
	key string,
) error {

	return r.cache.Delete(
		ctx,
		key,
	)
}
