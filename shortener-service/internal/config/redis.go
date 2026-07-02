package config

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type RedisConfig struct {
	Client *redis.Client
	Log    *logrus.Logger
}

func NewRedisClient(
	cfg RedisSettings,
	log *logrus.Logger,
) *RedisConfig {

	addr := fmt.Sprintf(
		"%s:%s",
		cfg.Host,
		cfg.Port,
	)

	client := redis.NewClient(
		&redis.Options{
			Addr:         addr,
			Password:     cfg.Password,
			DB:           cfg.DB,
			ReadTimeout:  cfg.ReadTimeout,
			WriteTimeout: cfg.WriteTimeout,
			PoolSize:     cfg.PoolSize,
			MinIdleConns: cfg.MinIdleConns,
		},
	)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		3*time.Second,
	)

	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {

		log.WithError(err).
			Fatal(
				"failed to connect to Redis",
			)
	}

	log.Infof(
		"connected to Redis at %s",
		addr,
	)

	return &RedisConfig{
		Client: client,
		Log:    log,
	}
}
