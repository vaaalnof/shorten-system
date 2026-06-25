package route

import (
	"shortener-service/internal/delivery/http"
	"shortener-service/internal/delivery/http/middleware"

	"github.com/gofiber/fiber/v2"
)

type Config struct {
	App *fiber.App

	AuthMiddleware *middleware.AuthMiddleware

	URLController    *http.URLController
	ReportController *http.ReportController
	QRController     *http.QRController
}

func (c *Config) Setup() {

	// =====================================================
	// PUBLIC REDIRECT
	// =====================================================

	c.App.Get(
		"/:short_code",
		middleware.RequestContext(),
		c.URLController.Redirect,
	)

	c.App.Post(
		"/:short_code/verify",
		middleware.RequestContext(),
		c.URLController.VerifyPassword,
	)

	// =====================================================
	// API V1
	// =====================================================

	api := c.App.Group(
		"/api/v1",
	)

	// =====================================================
	// URLS
	// =====================================================

	urls := api.Group(
		"/shorten/urls",
		middleware.RequestContext(),
		c.AuthMiddleware.Handle,
	)

	urls.Post(
		"",
		c.URLController.Create,
	)

	urls.Get(
		"",
		c.URLController.List,
	)

	urls.Get(
		"/:id",
		c.URLController.Get,
	)

	urls.Patch(
		"/:id",
		c.URLController.Update,
	)

	urls.Patch(
		"/:id/password",
		c.URLController.UpdatePassword,
	)

	urls.Delete(
		"/:id/password",
		c.URLController.RemovePassword,
	)

	urls.Delete(
		"/:id",
		c.URLController.Delete,
	)

	urls.Get(
		"/:id/qrcode",
		c.QRController.Generate,
	)

	// =====================================================
	// REPORTS
	// =====================================================

	reports := api.Group(
		"/shorten/reports",
		middleware.RequestContext(),
		c.AuthMiddleware.Handle,
	)

	reports.Get(
		"/:id/summary",
		c.ReportController.Summary,
	)

	reports.Get(
		"/:id/chart",
		c.ReportController.Chart,
	)

	reports.Get(
		"/:id/referrers",
		c.ReportController.Referrers,
	)

	reports.Get(
		"/:id/countries",
		c.ReportController.Countries,
	)

	reports.Get(
		"/:id/devices",
		c.ReportController.Devices,
	)

}
