package config

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type RedisClient struct {
	Client *redis.Client
	Log    *logrus.Logger
}

func NewRedisClient(
	cfg RedisSettings,
	log *logrus.Logger,
) *RedisClient {

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
			PoolSize:     cfg.PoolSize,
			MinIdleConns: cfg.MinIdleConns,
			ReadTimeout:  cfg.ReadTimeout,
			WriteTimeout: cfg.WriteTimeout,
		},
	)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		3*time.Second,
	)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {

		log.WithError(err).Fatal(
			"failed to connect to redis",
		)
	}

	log.Infof(
		"connected to redis at %s",
		addr,
	)

	return &RedisClient{
		Client: client,
		Log:    log,
	}
}
