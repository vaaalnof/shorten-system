package http

import (
	"auth-service/internal/delivery/http/middleware"
	"auth-service/internal/usecase/auth"

	"github.com/gofiber/fiber/v2"
)

type LogoutAllController struct {
	UseCase *auth.LogoutAllUseCase
}

func NewLogoutAllController(
	uc *auth.LogoutAllUseCase,
) *LogoutAllController {

	return &LogoutAllController{
		UseCase: uc,
	}
}

func (c *LogoutAllController) Handle(
	ctx *fiber.Ctx,
) error {

	// =====================================================
	// META
	// =====================================================

	meta := middleware.FromFiber(
		ctx,
	)

	requestCtx := middleware.WithMeta(
		ctx.UserContext(),
		meta,
	)

	// =====================================================
	// USECASE
	// =====================================================

	res, err := c.UseCase.Execute(
		requestCtx,
	)

	if err != nil {
		return err
	}

	// =====================================================
	// RESPONSE
	// =====================================================

	return ctx.Status(
		fiber.StatusOK,
	).JSON(
		res,
	)
}
