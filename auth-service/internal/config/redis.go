package config

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type RedisConfig struct {
	Client *redis.Client
	Log    *logrus.Logger
}

// NewRedisClient membuat koneksi redis baru
func NewRedisClient(v *viper.Viper, log *logrus.Logger) *RedisConfig {
	host := v.GetString("redis.host")
	port := v.GetString("redis.port")
	password := v.GetString("redis.password")
	db := v.GetInt("redis.db")

	addr := fmt.Sprintf("%s:%s", host, port)
	client := redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     password,
		DB:           db,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		PoolSize:     10,
		MinIdleConns: 2,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		log.WithError(err).Fatal("Failed to connect to Redis")
	} else {
		log.Infof("Connected to Redis at %s", addr)
	}

	return &RedisConfig{
		Client: client,
		Log:    log,
	}
}
