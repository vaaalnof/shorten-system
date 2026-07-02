package http

import (
	"auth-service/internal/delivery/http/middleware"
	"auth-service/internal/usecase/user"

	"github.com/gofiber/fiber/v2"
)

type ProfileController struct {
	UseCase *user.ProfileUseCase
}

func NewProfileController(
	uc *user.ProfileUseCase,
) *ProfileController {

	return &ProfileController{
		UseCase: uc,
	}
}

func (c *ProfileController) Handle(
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
