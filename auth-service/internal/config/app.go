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

type BootstrapConfig struct {
	DB       *Database
	Redis    *RedisConfig
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
}

func Bootstrap(cfg *BootstrapConfig) {

	// === Middleware (pure infra) ===
	cfg.App.Use(recover.New())
	cfg.App.Use(logger.New())
	cfg.App.Use(requestid.New())
	cfg.App.Use(cors.New())

	// === Check DB Connection (infra health) ===
	if err := cfg.DB.Master.Ping(); err != nil {
		cfg.Log.Fatalf("DB Master not reachable: %v", err)
	}
	if err := cfg.DB.Slave.Ping(); err != nil {
		cfg.Log.Fatalf("DB Slave not reachable: %v", err)
	}

	// === Check Redis Connection (infra health) ===
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := cfg.Redis.Client.Ping(ctx).Err(); err != nil {
		cfg.Log.Fatalf("Redis not reachable: %v", err)
	}

	// === Kafka basic config check (infra) ===
	//brokers := cfg.Config.GetStringSlice("kafka.brokers")
	//if len(brokers) == 0 {
	//	brokers = []string{"localhost:9092"}
	//}

	//saramaCfg := sarama.NewConfig()
	//saramaCfg.Producer.RequiredAcks = sarama.WaitForAll
	//saramaCfg.Producer.Return.Successes = true

	cfg.Log.Info("Infra bootstrap completed successfully")
}
