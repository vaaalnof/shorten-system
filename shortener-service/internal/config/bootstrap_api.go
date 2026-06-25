package config

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type BootstrapAPIConfig struct {
	DB    *Database
	Redis *RedisConfig
	NATS  *NATSConfig

	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
}

func BootstrapAPI(
	cfg *BootstrapAPIConfig,
) {

	// =====================================================
	// MIDDLEWARE
	// =====================================================

	cfg.App.Use(
		recover.New(),
	)

	cfg.App.Use(
		logger.New(),
	)

	cfg.App.Use(
		requestid.New(),
	)

	cfg.App.Use(
		cors.New(),
	)

	// =====================================================
	// DATABASE
	// =====================================================

	if err := cfg.DB.Master.Ping(); err != nil {

		cfg.Log.Fatalf(
			"db master not reachable: %v",
			err,
		)
	}

	if err := cfg.DB.Slave.Ping(); err != nil {

		cfg.Log.Fatalf(
			"db slave not reachable: %v",
			err,
		)
	}

	// =====================================================
	// REDIS
	// =====================================================

	ctx, cancel := context.WithTimeout(
		context.Background(),
		3*time.Second,
	)

	defer cancel()

	if err := cfg.Redis.Client.Ping(
		ctx,
	).Err(); err != nil {

		cfg.Log.Fatalf(
			"redis not reachable: %v",
			err,
		)
	}

	// =====================================================
	// NATS
	// =====================================================

	if cfg.NATS == nil ||
		cfg.NATS.Conn == nil {

		cfg.Log.Fatal(
			"nats not initialized",
		)
	}

	if !cfg.NATS.Conn.IsConnected() {

		cfg.Log.Fatal(
			"nats not connected",
		)
	}

	cfg.Log.Info(
		"api bootstrap completed successfully",
	)
}
