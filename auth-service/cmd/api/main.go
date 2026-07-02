package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"os"
	"os/signal"
	"syscall"
	"time"

	"auth-service/internal/config"
	"auth-service/internal/delivery/http"
	"auth-service/internal/delivery/http/route"
	cacheinfra "auth-service/internal/infra/cache"
	mailinfra "auth-service/internal/infra/mail"
	"auth-service/internal/repository"
	dbrepo "auth-service/internal/repository/db"
	"auth-service/internal/security"
	"auth-service/internal/usecase/auth"
	"auth-service/internal/usecase/user"
)

func main() {

	// =====================================================
	// BASE CONFIG
	// =====================================================

	v := config.NewViper()

	settings := config.NewSettings(
		v,
	)

	log := config.NewLogger(
		settings.Log,
	)

	db := config.NewDatabase(
		settings.Database,
		log,
	)

	redis := config.NewRedisClient(
		settings.Redis,
		log,
	)

	smtpConfig := config.NewSMTPConfig(
		settings.SMTP,
		log,
	)

	validate := config.NewValidator()

	app := config.NewFiber(
		settings.Web,
	)

	app.Use(
		recover.New(),
	)

	app.Use(
		logger.New(),
	)

	app.Use(
		requestid.New(),
	)

	app.Use(
		cors.New(
			cors.Config{
				AllowOrigins: settings.Web.CORS.AllowOrigins,

				AllowMethods: settings.Web.CORS.AllowMethods,

				AllowHeaders: settings.Web.CORS.AllowHeaders,

				AllowCredentials: settings.Web.CORS.AllowCredentials,
			},
		),
	)

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

	mailer := mailinfra.NewSMTPMailer(
		smtpConfig,
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
		settings.JWT.Secret,
		settings.JWT.AccessTokenTTL,
		settings.JWT.RefreshTokenTTL,
	)

	googleOAuth := security.NewGoogleOAuth(
		settings.Google.ClientID,
		settings.Google.ClientSecret,
		settings.Google.RedirectURL,
	)

	rateLimiter := security.NewRateLimiter(
		cache,
	)

	// =====================================================
	// AUTH BASE USECASE
	// =====================================================

	accessSessionUseCase := auth.NewAccessSessionUseCase(
		cache,
		jwtService,
	)

	refreshSessionUseCase := auth.NewRefreshSessionUseCase(
		jwtService,
	)

	createSession := auth.NewCreateSession(
		userSessionRepo,
		cache,
		refreshTokenHash,
		jwtService,
		settings.Cache.SessionTTL,
	)

	// =====================================================
	// AUTH USECASE
	// =====================================================

	registerUseCase := auth.NewRegisterUseCase(
		baseRepo,
		validate,
		userRepo,
		authProviderRepo,
		passwordHash,
		rateLimiter,
		settings.RateLimit.Register.MaxAttempts,
		settings.RateLimit.Register.Window,
	)

	loginUseCase := auth.NewLoginUseCase(
		validate,
		userRepo,
		authProviderRepo,
		createSession,
		passwordHash,
		rateLimiter,
		settings.RateLimit.Login.MaxAttempts,
		settings.RateLimit.Login.Window,
	)

	googleLoginUseCase := auth.NewGoogleLoginUseCase(
		cache,
		googleOAuth,
		settings.Cache.OAuthStateTTL,
	)

	googleCallbackUseCase := auth.NewGoogleCallbackUseCase(
		baseRepo,
		validate,
		userRepo,
		authProviderRepo,
		createSession,
		cache,
		googleOAuth,
	)

	validateTokenUseCase := auth.NewValidateTokenUseCase(
		accessSessionUseCase,
	)

	refreshTokenUseCase := auth.NewRefreshTokenUseCase(
		refreshSessionUseCase,
		userSessionRepo,
		cache,
		refreshTokenHash,
		settings.Cache.SessionTTL,
	)

	logoutUseCase := auth.NewLogoutUseCase(
		accessSessionUseCase,
		userSessionRepo,
		cache,
	)

	logoutAllUseCase := auth.NewLogoutAllUseCase(
		accessSessionUseCase,
		userSessionRepo,
		cache,
	)

	// =====================================================
	// USER BASE USECASE
	// =====================================================

	userUseCase := user.NewUseCase(
		cache,
		userSessionRepo,
		jwtService,
		settings.Cache.SessionTTL,
	)

	profileUseCase := user.NewProfileUseCase(
		userUseCase,
		userRepo,
	)

	sendEmailVerificationCodeUseCase :=
		user.NewSendVerificationCodeUseCase(
			userUseCase,
			userRepo,
			mailer,
			settings.Cache.EmailVerificationTTL,
			settings.Cache.EmailVerificationCooldownTTL,
		)

	verificationCodeUseCase :=
		user.NewVerificationCodeUseCase(
			userUseCase,
			userRepo,
			cache,
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

	googleLoginController := http.NewGoogleLoginController(
		googleLoginUseCase,
	)

	googleCallbackController := http.NewGoogleCallbackController(
		googleCallbackUseCase,
	)

	validateTokenController := http.NewValidateTokenController(
		validateTokenUseCase,
	)

	profileController := http.NewProfileController(
		profileUseCase,
	)

	refreshTokenController := http.NewRefreshTokenController(
		refreshTokenUseCase,
	)

	logoutController := http.NewLogoutController(
		logoutUseCase,
	)

	logoutAllController := http.NewLogoutAllController(
		logoutAllUseCase,
	)

	sendEmailVerificationCodeController :=
		http.NewSendVerificationCodeController(
			sendEmailVerificationCodeUseCase,
		)

	verificationCodeController :=
		http.NewVerificationCodeController(
			verificationCodeUseCase,
		)

	// =====================================================
	// ROUTE
	// =====================================================

	routeConfig := route.Config{
		App: app,

		// AUTH
		RegisterUserController:   registerUserController,
		LoginController:          loginController,
		GoogleLoginController:    googleLoginController,
		GoogleCallbackController: googleCallbackController,
		ValidateTokenController:  validateTokenController,
		RefreshTokenController:   refreshTokenController,
		LogoutController:         logoutController,
		LogoutAllController:      logoutAllController,

		// USER
		ProfileController:                   profileController,
		SendEmailVerificationCodeController: sendEmailVerificationCodeController,
		VerificationCodeController:          verificationCodeController,
	}

	routeConfig.Setup()

	// =====================================================
	// SERVER START
	// =====================================================

	addr := fmt.Sprintf(":%d", settings.Web.Port)

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
