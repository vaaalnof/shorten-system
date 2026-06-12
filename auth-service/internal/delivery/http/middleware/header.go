package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func getClientIP(
	c *fiber.Ctx,
) string {

	if ip := c.Get(
		"CF-Connecting-IP",
	); ip != "" {
		return ip
	}

	if ip := c.Get(
		"X-Real-IP",
	); ip != "" {
		return ip
	}

	if ip := c.Get(
		"X-Forwarded-For",
	); ip != "" {

		parts := strings.Split(
			ip,
			",",
		)

		return strings.TrimSpace(
			parts[0],
		)
	}

	return c.IP()
}

func RequestContext() fiber.Handler {

	return func(
		c *fiber.Ctx,
	) error {

		deviceID := strings.TrimSpace(
			c.Get("X-Device-ID"),
		)

		deviceType := strings.TrimSpace(
			c.Get("X-Device-Type"),
		)

		clientIP := getClientIP(
			c,
		)

		userAgent := c.Get(
			"User-Agent",
		)

		auth := strings.TrimSpace(
			c.Get("Authorization"),
		)

		refreshToken := strings.TrimSpace(
			c.Get("X-Refresh-Token"),
		)

		c.Locals(
			"device_id",
			deviceID,
		)

		c.Locals(
			"device_type",
			deviceType,
		)

		c.Locals(
			"client_ip",
			clientIP,
		)

		c.Locals(
			"user_agent",
			userAgent,
		)

		c.Locals(
			"authorization",
			auth,
		)

		c.Locals(
			"refresh_token",
			refreshToken,
		)

		meta := &RequestMeta{
			DeviceID:     deviceID,
			DeviceType:   deviceType,
			ClientIP:     clientIP,
			UserAgent:    userAgent,
			Auth:         auth,
			RefreshToken: refreshToken,
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
