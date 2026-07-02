package route

import (
	"auth-service/internal/delivery/http"
	"auth-service/internal/delivery/http/middleware"

	"github.com/gofiber/fiber/v2"
)

type Config struct {
	App *fiber.App

	// AUTH
	RegisterUserController   *http.RegisterUserController
	LoginController          *http.LoginController
	GoogleLoginController    *http.GoogleLoginController
	GoogleCallbackController *http.GoogleCallbackController
	ValidateTokenController  *http.ValidateTokenController
	RefreshTokenController   *http.RefreshTokenController
	LogoutController         *http.LogoutController
	LogoutAllController      *http.LogoutAllController

	// USER
	ProfileController                   *http.ProfileController
	SendEmailVerificationCodeController *http.SendVerificationCodeController
	VerificationCodeController          *http.VerificationCodeController
}

func (c *Config) Setup() {

	api := c.App.Group(
		"/api/v1",
	)

	auth := api.Group(
		"/auth",
	)

	user := api.Group(
		"/user",
	)

	c.setupAuthRoute(
		auth,
	)

	c.setupUserRoute(
		user,
	)
}

// =====================================================
// AUTH ROUTES
// =====================================================

func (c *Config) setupAuthRoute(
	auth fiber.Router,
) {

	if c.RegisterUserController != nil {

		auth.Post(
			"/register",
			middleware.RequestContext(),
			c.RegisterUserController.Handle,
		)
	}

	if c.LoginController != nil {

		auth.Post(
			"/login",
			middleware.RequestContext(),
			c.LoginController.Handle,
		)
	}

	// =====================================================
	// GOOGLE LOGIN
	// =====================================================
	if c.GoogleLoginController != nil {
		auth.Get(
			"/google",
			middleware.RequestContext(),
			c.GoogleLoginController.Handle,
		)
	}

	if c.GoogleCallbackController != nil {

		auth.Get(
			"/google/callback",
			middleware.RequestContext(),
			c.GoogleCallbackController.Handle,
		)

	}

	if c.ValidateTokenController != nil {

		auth.Get(
			"/validate-token",
			middleware.RequestContext(),
			c.ValidateTokenController.Handle,
		)
	}

	if c.RefreshTokenController != nil {

		auth.Post(
			"/refresh-token",
			middleware.RequestContext(),
			c.RefreshTokenController.Handle,
		)
	}

	if c.LogoutController != nil {

		auth.Post(
			"/logout",
			middleware.RequestContext(),
			c.LogoutController.Handle,
		)
	}

	if c.LogoutAllController != nil {
		auth.Post(
			"/logout-all",
			middleware.RequestContext(),
			c.LogoutAllController.Handle,
		)
	}
}

// =====================================================
// USER ROUTES
// =====================================================

func (c *Config) setupUserRoute(
	user fiber.Router,
) {

	if c.ProfileController != nil {

		user.Get(
			"/profile",
			middleware.RequestContext(),
			c.ProfileController.Handle,
		)
	}

	if c.SendEmailVerificationCodeController != nil {

		user.Post(
			"/email-verification/send",
			middleware.RequestContext(),
			c.SendEmailVerificationCodeController.Handle,
		)
	}

	if c.VerificationCodeController != nil {

		user.Post(
			"/email-verification/verify",
			middleware.RequestContext(),
			c.VerificationCodeController.Handle,
		)
	}
}
