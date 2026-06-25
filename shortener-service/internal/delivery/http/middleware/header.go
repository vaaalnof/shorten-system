package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func getClientIP(
	c *fiber.Ctx,
) string {

	if ip := strings.TrimSpace(
		c.Get("CF-Connecting-IP"),
	); ip != "" {

		return ip
	}

	if ip := strings.TrimSpace(
		c.Get("X-Forwarded-For"),
	); ip != "" {

		parts := strings.Split(
			ip,
			",",
		)

		return strings.TrimSpace(
			parts[0],
		)
	}

	if ip := strings.TrimSpace(
		c.Get("X-Real-IP"),
	); ip != "" {

		return ip
	}

	return c.IP()
}

func RequestContext() fiber.Handler {

	return func(
		c *fiber.Ctx,
	) error {

		auth := strings.TrimSpace(
			c.Get("Authorization"),
		)

		referer := strings.TrimSpace(
			c.Get("Referer"),
		)

		userAgent := strings.TrimSpace(
			c.Get("User-Agent"),
		)

		ipAddress := getClientIP(
			c,
		)

		c.Locals(
			"authorization",
			auth,
		)

		c.Locals(
			"referer",
			referer,
		)

		c.Locals(
			"user_agent",
			userAgent,
		)

		c.Locals(
			"ip_address",
			ipAddress,
		)

		meta := &RequestMeta{
			Auth:      auth,
			Referer:   referer,
			UserAgent: userAgent,
			IPAddress: ipAddress,
		}

		ctx := WithMeta(
			c.UserContext(),
			meta,
		)

		c.SetUserContext(
			ctx,
		)

		return c.Next()
	}
}
