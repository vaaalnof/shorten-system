package http

import (
	"auth-service/internal/delivery/http/middleware"
	"auth-service/internal/model"
	"auth-service/internal/usecase/user"

	"github.com/gofiber/fiber/v2"
)

type VerificationCodeController struct {
	UseCase *user.VerificationCodeUseCase
}

func NewVerificationCodeController(
	uc *user.VerificationCodeUseCase,
) *VerificationCodeController {

	return &VerificationCodeController{
		UseCase: uc,
	}
}

func (c *VerificationCodeController) Handle(
	ctx *fiber.Ctx,
) error {

	// =====================================================
	// REQUEST
	// =====================================================

	request := new(
		model.VerificationCodeRequest,
	)

	if err := ctx.BodyParser(
		request,
	); err != nil {

		return err
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
		request,
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
