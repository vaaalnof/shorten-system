package http

import (
	"auth-service/internal/delivery/http/middleware"
	"auth-service/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type RefreshTokenController struct {
	UseCase *usecase.RefreshTokenUseCase
}

func NewRefreshTokenController(
	uc *usecase.RefreshTokenUseCase,
) *RefreshTokenController {
	return &RefreshTokenController{
		UseCase: uc,
	}
}

func (c *RefreshTokenController) Handle(
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
