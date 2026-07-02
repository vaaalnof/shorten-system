package middleware

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

type RequestMeta struct {
	ClientIP     string
	UserAgent    string
	Auth         string
	RefreshToken string

	// auth middleware
	UserID    string
	SessionID string
}

type key string

const requestKey key = "request_meta"

func getLocalString(
	c *fiber.Ctx,
	k string,
) string {

	if v, ok := c.Locals(
		k,
	).(string); ok {

		return v
	}

	return ""
}

func FromFiber(
	c *fiber.Ctx,
) *RequestMeta {

	return &RequestMeta{
		ClientIP: getLocalString(
			c,
			"client_ip",
		),
		UserAgent: getLocalString(
			c,
			"user_agent",
		),
		Auth: getLocalString(
			c,
			"authorization",
		),
		RefreshToken: getLocalString(
			c,
			"refresh_token",
		),
		UserID: getLocalString(
			c,
			"user_id",
		),
		SessionID: getLocalString(
			c,
			"session_id",
		),
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
