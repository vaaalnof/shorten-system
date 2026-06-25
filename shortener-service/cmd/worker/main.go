package main

import (
	"context"
	"os"
	"os/signal"
	"shortener-service/internal/usecase/analytics"
	"syscall"

	"shortener-service/internal/config"
	natsinfra "shortener-service/internal/infra/nats"
	"shortener-service/internal/repository"
	dbrepo "shortener-service/internal/repository/db"
	"shortener-service/internal/security"
	"shortener-service/internal/worker"
)

func main() {

	// =====================================================
	// CONFIG
	// =====================================================

	viperConfig := config.NewViper()

	settings := config.NewSettings(
		viperConfig,
	)

	log := config.NewLogger(
		viperConfig,
	)

	db := config.NewDatabase(
		viperConfig,
		log,
	)

	natsConfig := config.NewNATSClient(
		viperConfig,
		log,
	)

	// =====================================================
	// BOOTSTRAP
	// =====================================================

	config.BootstrapWorker(
		&config.BootstrapWorkerConfig{
			DB:     db,
			NATS:   natsConfig,
			Log:    log,
			Config: viperConfig,
		},
	)

	// =====================================================
	// NATS STREAM SETUP
	// =====================================================

	if err := natsinfra.SetupAnalytics(
		natsConfig.JS,
		settings,
	); err != nil {

		log.Fatalf(
			"failed to setup analytics stream: %v",
			err,
		)
	}

	// =====================================================
	// REPOSITORY
	// =====================================================

	baseRepo := repository.NewRepository(
		db.Master,
		db.Slave,
		log,
	)

	analyticsEventRepo := dbrepo.NewAnalyticsEventRepo(
		baseRepo,
	)

	urlDailyAnalyticsRepo := dbrepo.NewURLDailyAnalyticsRepo(
		baseRepo,
	)

	urlDailyVisitorRepo := dbrepo.NewURLDailyVisitorRepo(
		baseRepo,
	)

	// =====================================================
	// SECURITY
	// =====================================================

	visitorHash := security.NewVisitorHash()

	geoIP, err := security.NewGeoIP(
		settings.GeoIPDatabasePath,
	)

	if err != nil {

		log.Fatalf(
			"failed to load geoip database: %v",
			err,
		)
	}

	log.WithField(
		"path",
		settings.GeoIPDatabasePath,
	).Info(
		"geoip database loaded",
	)

	// =====================================================
	// CLEANUP
	// =====================================================

	defer func() {

		_ = geoIP.Close()

		natsConfig.Conn.Close()

		_ = db.Master.Close()

		_ = db.Slave.Close()
	}()

	// =====================================================
	// USECASE
	// =====================================================

	analyticsProcessor := analytics.NewAnalyticsEventProcessor(
		analyticsEventRepo,
		urlDailyAnalyticsRepo,
		urlDailyVisitorRepo,
		visitorHash,
		geoIP,
	)

	// =====================================================
	// NATS CONSUMER
	// =====================================================

	analyticsConsumer := natsinfra.NewConsumer(
		natsConfig.JS,
		settings,
		log,
	)

	// =====================================================
	// WORKER
	// =====================================================

	analyticsWorker := worker.NewAnalyticsWorker(
		analyticsConsumer,
		analyticsProcessor,
	)

	ctx, cancel := context.WithCancel(
		context.Background(),
	)

	defer cancel()

	go func() {

		if err := analyticsWorker.Run(
			ctx,
		); err != nil {

			log.Errorf(
				"analytics worker stopped: %v",
				err,
			)

			cancel()
		}
	}()

	log.Info(
		"analytics worker started",
	)

	// =====================================================
	// WAIT SIGNAL
	// =====================================================

	quit := make(
		chan os.Signal,
		1,
	)

	signal.Notify(
		quit,
		os.Interrupt,
		syscall.SIGTERM,
	)

	<-quit

	log.Info(
		"shutting down worker...",
	)

	cancel()

	log.Info(
		"worker stopped gracefully",
	)
}
