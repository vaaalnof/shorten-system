package cache

import (
	"context"
	"time"

	"auth-service/internal/usecase/port"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

var _ port.Cache = (*RedisCache)(nil)

type RedisCache struct {
	Client *redis.Client
	Log    *logrus.Logger
}

func NewRedisCache(
	client *redis.Client,
	log *logrus.Logger,
) *RedisCache {
	return &RedisCache{
		Client: client,
		Log:    log,
	}
}

func (r *RedisCache) Set(
	ctx context.Context,
	key string,
	value interface{},
	ttl time.Duration,
) error {

	if err := r.Client.Set(
		ctx,
		key,
		value,
		ttl,
	).Err(); err != nil {

		r.Log.WithError(err).
			WithField("key", key).
			Error("redis set failed")

		return err
	}

	return nil
}

func (r *RedisCache) Get(
	ctx context.Context,
	key string,
) (string, error) {

	val, err := r.Client.Get(
		ctx,
		key,
	).Result()

	if err == redis.Nil {
		return "", nil
	}

	if err != nil {

		r.Log.WithError(err).
			WithField("key", key).
			Error("redis get failed")

		return "", err
	}

	return val, nil
}

func (r *RedisCache) Delete(
	ctx context.Context,
	keys ...string,
) error {

	if err := r.Client.Del(
		ctx,
		keys...,
	).Err(); err != nil {

		r.Log.WithError(err).
			WithField("keys", keys).
			Error("redis delete failed")

		return err
	}

	return nil
}

func (r *RedisCache) Exists(
	ctx context.Context,
	key string,
) bool {

	count, err := r.Client.Exists(
		ctx,
		key,
	).Result()

	if err != nil {

		r.Log.WithError(err).
			WithField("key", key).
			Warn("redis exists failed")

		return false
	}

	return count > 0
}

func (r *RedisCache) Incr(
	ctx context.Context,
	key string,
	ttl time.Duration,
) (int64, error) {

	val, err := r.Client.Incr(
		ctx,
		key,
	).Result()

	if err != nil {

		r.Log.WithError(err).
			WithField("key", key).
			Error("redis incr failed")

		return 0, err
	}

	if val == 1 {

		if err := r.Client.Expire(
			ctx,
			key,
			ttl,
		).Err(); err != nil {

			r.Log.WithError(err).
				WithField("key", key).
				Error("redis expire failed")

			return 0, err
		}
	}

	return val, nil
}

func (r *RedisCache) Expire(
	ctx context.Context,
	key string,
	ttl time.Duration,
) error {

	if err := r.Client.Expire(
		ctx,
		key,
		ttl,
	).Err(); err != nil {

		r.Log.WithError(err).
			WithField("key", key).
			Error("redis expire failed")

		return err
	}

	return nil
}

func (r *RedisCache) TTL(
	ctx context.Context,
	key string,
) (time.Duration, error) {

	ttl, err := r.Client.TTL(
		ctx,
		key,
	).Result()

	if err != nil {

		r.Log.WithError(err).
			WithField("key", key).
			Error("redis ttl failed")

		return 0, err
	}

	return ttl, nil
}

func (r *RedisCache) FlushAll(
	ctx context.Context,
) error {

	if err := r.Client.FlushAll(
		ctx,
	).Err(); err != nil {

		r.Log.WithError(err).
			Error("redis flushall failed")

		return err
	}

	return nil
}
