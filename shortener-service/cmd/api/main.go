package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"shortener-service/internal/usecase/qr"
	"shortener-service/internal/usecase/report"
	"syscall"
	"time"

	"shortener-service/internal/config"
	"shortener-service/internal/delivery/http"
	"shortener-service/internal/delivery/http/middleware"
	"shortener-service/internal/delivery/http/route"
	"shortener-service/internal/usecase/analytics"
	"shortener-service/internal/usecase/shorturl"

	authinfra "shortener-service/internal/infra/auth"
	cacheinfra "shortener-service/internal/infra/cache"
	natsinfra "shortener-service/internal/infra/nats"

	"shortener-service/internal/repository"
	dbrepo "shortener-service/internal/repository/db"

	"shortener-service/internal/security"
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

	redis := config.NewRedisClient(
		viperConfig,
		log,
	)

	natsConfig := config.NewNATSClient(
		viperConfig,
		log,
	)

	validate := config.NewValidator()

	app := config.NewFiber(
		viperConfig,
	)

	// =====================================================
	// BOOTSTRAP
	// =====================================================

	config.BootstrapAPI(
		&config.BootstrapAPIConfig{
			DB:       db,
			Redis:    redis,
			NATS:     natsConfig,
			App:      app,
			Log:      log,
			Validate: validate,
			Config:   viperConfig,
		},
	)

	// =====================================================
	// REPOSITORY
	// =====================================================

	baseRepo := repository.NewRepository(

		db.Master,
		db.Slave,
		log,
	)

	redisCache := cacheinfra.NewRedisCache(

		redis.Client,
		log,
	)

	urlRepo := dbrepo.NewURLRepo(

		baseRepo,
	)

	reportRepo := dbrepo.NewReportRepo(

		baseRepo,
	)

	reservedAliasRepo := dbrepo.NewReservedAliasRepo(

		baseRepo,
	)

	// =====================================================
	// NATS
	// =====================================================

	analyticsProducer := natsinfra.NewProducer(
		natsConfig.JS,
		log,
	)

	// =====================================================
	// ANALYTICS
	// =====================================================

	analyticsPublisher := analytics.NewAnalyticsEventPublisher(

		analyticsProducer,
		log,
	)

	// =====================================================
	// AUTH
	// =====================================================

	authClient := authinfra.NewClient(
		settings.AuthServiceBaseURL,
		settings.AuthServiceTimeout,
	)

	// =====================================================
	// SECURITY
	// =====================================================

	passwordHash := security.NewPasswordHash()

	// =====================================================
	// MIDDLEWARE
	// =====================================================

	authMiddleware := middleware.NewAuthMiddleware(
		authClient,
	)

	// =====================================================
	// USECASE
	// =====================================================

	urlUseCase := shorturl.NewURLUseCase(
		validate,
		urlRepo,
		reservedAliasRepo,
		passwordHash,
		redisCache,
		settings.URLCacheTTL,
		analyticsPublisher,
		log,
	)

	reportUseCase := report.NewReportUseCase(
		validate,
		urlRepo,
		reportRepo,
		log,
	)

	qrUseCase := qr.NewQRUseCase(
		validate,
		urlRepo,
		settings.ShortenerBaseURL,
		log,
	)

	// =====================================================
	// CONTROLLER
	// =====================================================

	urlController := http.NewURLController(
		urlUseCase,
	)

	reportController := http.NewReportController(
		reportUseCase,
	)

	qrController := http.NewQRController(
		qrUseCase,
	)

	// =====================================================
	// ROUTE
	// =====================================================

	routeConfig := route.Config{
		App: app,

		AuthMiddleware: authMiddleware,

		URLController:    urlController,
		ReportController: reportController,
		QRController:     qrController,
	}

	routeConfig.Setup()

	// =====================================================
	// START SERVER
	// =====================================================

	addr := fmt.Sprintf(
		":%d",
		viperConfig.GetInt(
			"web.port",
		),
	)

	go func() {

		if err := app.Listen(
			addr,
		); err != nil {

			log.Errorf(
				"fiber stopped: %v",
				err,
			)
		}
	}()

	log.Infof(
		"server running on %s",
		addr,
	)

	// =====================================================
	// SHUTDOWN SIGNAL
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
		"shutting down server...",
	)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)

	defer cancel()

	done := make(
		chan struct{},
	)

	go func() {

		_ = app.Shutdown()

		close(done)
	}()

	select {

	case <-done:

	case <-ctx.Done():

		log.Warn(
			"force shutdown...",
		)
	}

	// =====================================================
	// CLOSE NATS
	// =====================================================

	if natsConfig != nil &&
		natsConfig.Conn != nil {

		natsConfig.Conn.Close()
	}

	// =====================================================
	// CLOSE REDIS
	// =====================================================

	if redis != nil &&
		redis.Client != nil {

		_ = redis.Client.Close()
	}

	// =====================================================
	// CLOSE DATABASE
	// =====================================================

	if db != nil {

		if db.Master != nil {
			_ = db.Master.Close()
		}

		if db.Slave != nil {
			_ = db.Slave.Close()
		}
	}

	log.Info(
		"server stopped gracefully",
	)
}
