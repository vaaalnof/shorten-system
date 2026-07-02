package http

import (
	"auth-service/internal/delivery/http/middleware"
	"auth-service/internal/usecase/auth"

	"github.com/gofiber/fiber/v2"
)

type GoogleLoginController struct {
	UseCase *auth.GoogleLoginUseCase
}

func NewGoogleLoginController(
	uc *auth.GoogleLoginUseCase,
) *GoogleLoginController {

	return &GoogleLoginController{
		UseCase: uc,
	}
}

func (c *GoogleLoginController) Handle(
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
