package http

import (
	"auth-service/internal/delivery/http/middleware"
	"auth-service/internal/model"
	"auth-service/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type LoginController struct {
	UseCase *usecase.LoginUseCase
}

func NewLoginController(
	uc *usecase.LoginUseCase,
) *LoginController {
	return &LoginController{
		UseCase: uc,
	}
}

func (c *LoginController) Handle(
	ctx *fiber.Ctx,
) error {

	// =====================================================
	// REQUEST
	// =====================================================

	var req model.LoginRequest

	if err := ctx.BodyParser(
		&req,
	); err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"invalid request body",
		)
	}

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
		&req,
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
