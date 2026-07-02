package http

import (
	"auth-service/internal/delivery/http/middleware"
	"auth-service/internal/model"
	"auth-service/internal/usecase/auth"

	"github.com/gofiber/fiber/v2"
)

type RegisterUserController struct {
	UseCase *auth.RegisterUseCase
}

func NewRegisterUserController(
	uc *auth.RegisterUseCase,
) *RegisterUserController {
	return &RegisterUserController{
		UseCase: uc,
	}
}

func (c *RegisterUserController) Handle(
	ctx *fiber.Ctx,
) error {

	req := new(model.RegisterUserRequest)

	if err := ctx.BodyParser(req); err != nil {
		return fiber.ErrBadRequest
	}

	meta := middleware.FromFiber(
		ctx,
	)

	newCtx := middleware.WithMeta(
		ctx.UserContext(),
		meta,
	)

	res, err := c.UseCase.Execute(
		newCtx,
		req,
	)
	if err != nil {
		return err
	}

	return ctx.Status(
		fiber.StatusCreated,
	).JSON(
		res,
	)
}
