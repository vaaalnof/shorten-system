package http

import (
	"auth-service/internal/delivery/http/middleware"
	"auth-service/internal/usecase/user"
	"github.com/gofiber/fiber/v2"
)

type SendVerificationCodeController struct {
	UseCase *user.SendVerificationCodeUseCase
}

func NewSendVerificationCodeController(
	uc *user.SendVerificationCodeUseCase,
) *SendVerificationCodeController {

	return &SendVerificationCodeController{
		UseCase: uc,
	}
}

func (c *SendVerificationCodeController) Handle(
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
