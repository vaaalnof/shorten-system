package middleware

import (
	authinfra "shortener-service/internal/infra/auth"

	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware struct {
	authClient *authinfra.Client
}

func NewAuthMiddleware(
	authClient *authinfra.Client,
) *AuthMiddleware {

	return &AuthMiddleware{
		authClient: authClient,
	}
}

func (m *AuthMiddleware) Handle(
	c *fiber.Ctx,
) error {

	meta := GetMeta(
		c.UserContext(),
	)

	if meta == nil ||
		meta.Auth == "" {

		return fiber.NewError(
			fiber.StatusUnauthorized,
			"missing authorization token",
		)
	}

	userID,
		sessionID,
		emailVerified,
		err := m.authClient.ValidateToken(
		c.UserContext(),
		meta.Auth,
	)

	if err != nil {

		return fiber.NewError(
			fiber.StatusUnauthorized,
			"invalid token",
		)
	}

	meta.UserID = userID
	meta.SessionID = sessionID
	meta.EmailVerified = emailVerified

	ctx := WithMeta(
		c.UserContext(),
		meta,
	)

	c.SetUserContext(
		ctx,
	)

	return c.Next()
}
