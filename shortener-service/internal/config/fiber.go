package config

import (
	"shortener-service/internal/exception"
	"shortener-service/internal/model"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func NewFiber(
	config *viper.Viper,
) *fiber.App {

	app := fiber.New(
		fiber.Config{
			AppName: config.GetString(
				"app.name",
			),
			ErrorHandler: NewErrorHandler(),
			Prefork: config.GetBool(
				"web.prefork",
			),
		},
	)

	return app
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
