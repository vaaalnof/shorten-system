package middleware

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

type RequestMeta struct {
	DeviceID     string
	DeviceType   string
	ClientIP     string
	UserAgent    string
	Auth         string
	RefreshToken string
}

type key string

const requestKey key = "request_meta"

func getLocalString(c *fiber.Ctx, k string) string {
	if v, ok := c.Locals(k).(string); ok {
		return v
	}
	return ""
}

func FromFiber(c *fiber.Ctx) *RequestMeta {
	return &RequestMeta{
		DeviceID:     getLocalString(c, "device_id"),
		DeviceType:   getLocalString(c, "device_type"),
		ClientIP:     getLocalString(c, "client_ip"),
		UserAgent:    getLocalString(c, "user_agent"),
		Auth:         getLocalString(c, "authorization"),
		RefreshToken: getLocalString(c, "refresh_token"),
	}
}

func WithMeta(
	ctx context.Context,
	meta *RequestMeta,
) context.Context {
	return context.WithValue(
		ctx,
		requestKey,
		meta,
	)
}

func GetMeta(
	ctx context.Context,
) *RequestMeta {
	if v, ok := ctx.Value(
		requestKey,
	).(*RequestMeta); ok {
		return v
	}

	return nil
}
