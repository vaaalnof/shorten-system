package http

import (
	"auth-service/internal/delivery/http/middleware"
	"auth-service/internal/model"
	"auth-service/internal/usecase/auth"

	"github.com/gofiber/fiber/v2"
)

type GoogleCallbackController struct {
	UseCase *auth.GoogleCallbackUseCase
}

func NewGoogleCallbackController(
	uc *auth.GoogleCallbackUseCase,
) *GoogleCallbackController {

	return &GoogleCallbackController{
		UseCase: uc,
	}
}

func (c *GoogleCallbackController) Handle(
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
	// REQUEST
	// =====================================================

	req := &model.GoogleCallbackRequest{
		Code: ctx.Query(
			"code",
		),
		State: ctx.Query(
			"state",
		),
	}

	// =====================================================
	// USECASE
	// =====================================================

	res, err := c.UseCase.Execute(
		requestCtx,
		req,
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
