package main

import (
	"auth-service/internal/config"
	"auth-service/internal/delivery/http"
	"auth-service/internal/delivery/http/route"
	cacheinfra "auth-service/internal/infra/cache"
	"auth-service/internal/repository"
	dbrepo "auth-service/internal/repository/db"
	"auth-service/internal/security"
	"auth-service/internal/usecase"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	// =====================================================
	// BASE CONFIG
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

	validate := config.NewValidator()

	app := config.NewFiber(
		viperConfig,
	)

	// =====================================================
	// BOOTSTRAP
	// =====================================================

	config.Bootstrap(&config.BootstrapConfig{
		DB:       db,
		Redis:    redis,
		App:      app,
		Log:      log,
		Validate: validate,
		Config:   viperConfig,
	})

	// =====================================================
	// REPOSITORY
	// =====================================================

	baseRepo := repository.NewRepository(
		db.Master,
		db.Slave,
		log,
	)

	cache := cacheinfra.NewRedisCache(
		redis.Client,
		log,
	)

	userRepo := dbrepo.NewUserRepo(
		baseRepo,
	)

	authProviderRepo := dbrepo.NewUserAuthProviderRepo(
		baseRepo,
	)

	userSessionRepo := dbrepo.NewUserSessionRepo(
		baseRepo,
	)

	// =====================================================
	// SECURITY
	// =====================================================

	passwordHash := security.NewBcryptHash()

	refreshTokenHash := security.NewRefreshTokenHash()

	jwtService := security.NewJWTService(
		settings.JWTSecret,
		settings.AccessTokenTTL,
		settings.RefreshTokenTTL,
	)

	rateLimiter := security.NewRateLimiter(
		cache,
	)

	// =====================================================
	// USECASE
	// =====================================================

	registerUseCase := usecase.NewRegisterUseCase(
		baseRepo,
		validate,
		userRepo,
		authProviderRepo,
		passwordHash,
		rateLimiter,
		settings.RegisterMaxAttempts,
		settings.RegisterWindowTTL,
	)

	loginUseCase := usecase.NewLoginUseCase(
		validate,
		userRepo,
		authProviderRepo,
		userSessionRepo,
		cache,
		passwordHash,
		refreshTokenHash,
		jwtService,
		rateLimiter,
		settings.LoginMaxAttempts,
		settings.LoginWindowTTL,
		settings.SessionTTL,
	)

	validateTokenUseCase := usecase.NewValidateTokenUseCase(
		cache,
		jwtService,
	)

	meUseCase := usecase.NewMeUseCase(
		userRepo,
		userSessionRepo,
		cache,
		jwtService,
		settings.SessionTTL,
	)

	refreshTokenUseCase := usecase.NewRefreshTokenUseCase(
		userSessionRepo,
		cache,
		jwtService,
		refreshTokenHash,
		settings.SessionTTL,
	)

	logoutUseCase := usecase.NewLogoutUseCase(
		userSessionRepo,
		cache,
		jwtService,
	)

	// =====================================================
	// CONTROLLER
	// =====================================================

	registerUserController := http.NewRegisterUserController(
		registerUseCase,
	)

	loginController := http.NewLoginController(
		loginUseCase,
	)

	validateTokenController := http.NewValidateTokenController(
		validateTokenUseCase,
	)

	meController := http.NewMeController(
		meUseCase,
	)

	refreshTokenController := http.NewRefreshTokenController(
		refreshTokenUseCase,
	)

	logoutController := http.NewLogoutController(
		logoutUseCase,
	)

	// =====================================================
	// ROUTE
	// =====================================================

	routeConfig := route.Config{
		App: app,

		RegisterUserController:  registerUserController,
		LoginController:         loginController,
		ValidateTokenController: validateTokenController,
		MeController:            meController,
		RefreshTokenController:  refreshTokenController,
		LogoutController:        logoutController,
	}

	routeConfig.Setup()

	// =====================================================
	// KAFKA & WORKER
	// DISABLED FOR NOW
	// =====================================================

	// kafkaProducer := config.NewKafkaProducer(
	// 	viperConfig,
	// 	log,
	// )

	// registerConsumer := worker.NewRegisterConsumer(
	// 	kafkaProducer,
	// 	log,
	// )

	// go registerConsumer.Start()

	// =====================================================
	// SERVER START
	// =====================================================

	webPort := viperConfig.GetInt(
		"web.port",
	)

	addr := fmt.Sprintf(
		":%d",
		webPort,
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
	// GRACEFUL SHUTDOWN
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
