package config

import (
	"auth-service/internal/exception"
	"auth-service/internal/model"

	"github.com/gofiber/fiber/v2"
)

func NewFiber(
	cfg WebSettings,
) *fiber.App {

	return fiber.New(
		fiber.Config{
			AppName: cfg.AppName,
			Prefork: cfg.Prefork,

			ErrorHandler: NewErrorHandler(),
		},
	)
}

func NewErrorHandler() fiber.ErrorHandler {

	return func(
		ctx *fiber.Ctx,
		err error,
	) error {

		// =====================================================
		// APP ERROR
		// =====================================================

		if appErr, ok := err.(*exception.AppError); ok {

			return ctx.Status(
				appErr.StatusCode,
			).JSON(
				model.WebResponse[any]{
					Message: appErr.Message,
					Errors:  appErr.Errors,
				},
			)
		}

		// =====================================================
		// FIBER ERROR
		// =====================================================

		if fiberErr, ok := err.(*fiber.Error); ok {

			return ctx.Status(
				fiberErr.Code,
			).JSON(
				model.WebResponse[any]{
					Message: fiberErr.Message,
				},
			)
		}

		// =====================================================
		// INTERNAL SERVER ERROR
		// =====================================================

		return ctx.Status(
			fiber.StatusInternalServerError,
		).JSON(
			model.WebResponse[any]{
				Message: "internal server error",
			},
		)
	}
}
