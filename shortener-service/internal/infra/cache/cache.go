package cache

import (
	"context"
	"time"

	"shortener-service/internal/usecase/port"

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
	value string,
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
