package route

import (
	"auth-service/internal/delivery/http"
	"auth-service/internal/delivery/http/middleware"

	"github.com/gofiber/fiber/v2"
)

type Config struct {
	App *fiber.App

	// auth
	RegisterUserController  *http.RegisterUserController
	LoginController         *http.LoginController
	ValidateTokenController *http.ValidateTokenController
	MeController            *http.MeController
	RefreshTokenController  *http.RefreshTokenController
	LogoutController        *http.LogoutController
}

func (c *Config) Setup() {

	api := c.App.Group(
		"/api",
	)

	c.setupAuthRoute(
		api,
	)
}

// ================= AUTH =================
func (c *Config) setupAuthRoute(
	api fiber.Router,
) {

	// =====================================================
	// REGISTER
	// =====================================================

	if c.RegisterUserController != nil {

		api.Post(
			"/register",
			middleware.RequestContext(),
			c.RegisterUserController.Handle,
		)
	}

	// =====================================================
	// LOGIN
	// =====================================================

	if c.LoginController != nil {

		api.Post(
			"/login",
			middleware.RequestContext(),
			c.LoginController.Handle,
		)
	}

	// =====================================================
	// VALIDATE
	// =====================================================

	if c.ValidateTokenController != nil {
		api.Get(
			"/validate-token",
			middleware.RequestContext(),
			c.ValidateTokenController.Handle,
		)
	}

	// =====================================================
	// ME
	// =====================================================

	if c.MeController != nil {

		api.Get(
			"/me",
			middleware.RequestContext(),
			c.MeController.Handle,
		)
	}

	// =====================================================
	// REFRESH TOKEN
	// =====================================================

	if c.RefreshTokenController != nil {

		api.Post(
			"/refresh-token",
			middleware.RequestContext(),
			c.RefreshTokenController.Handle,
		)
	}

	// =====================================================
	// REFRESH TOKEN
	// =====================================================

	if c.LogoutController != nil {

		api.Post(
			"/logout",
			middleware.RequestContext(),
			c.LogoutController.Handle,
		)
	}
}
