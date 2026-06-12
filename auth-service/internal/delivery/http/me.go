package http

import (
	"auth-service/internal/delivery/http/middleware"
	"auth-service/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type MeController struct {
	UseCase *usecase.MeUseCase
}

func NewMeController(
	uc *usecase.MeUseCase,
) *MeController {
	return &MeController{
		UseCase: uc,
	}
}

func (c *MeController) Handle(
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
