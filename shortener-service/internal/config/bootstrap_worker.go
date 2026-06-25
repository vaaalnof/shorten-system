package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type BootstrapWorkerConfig struct {
	DB   *Database
	NATS *NATSConfig

	Log    *logrus.Logger
	Config *viper.Viper
}

func BootstrapWorker(
	cfg *BootstrapWorkerConfig,
) {

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
		"worker bootstrap completed successfully",
	)
}
