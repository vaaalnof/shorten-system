package config

import (
	"errors"

	"shortener-service/internal/exception"
	"shortener-service/internal/model"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewFiber(
	config *viper.Viper,
) *fiber.App {

	return fiber.New(
		fiber.Config{
			AppName: config.GetString(
				"app.name",
			),
			Prefork: config.GetBool(
				"web.prefork",
			),
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

		var appErr *exception.AppError

		if errors.As(
			err,
			&appErr,
		) {

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

		var fiberErr *fiber.Error

		if errors.As(
			err,
			&fiberErr,
		) {

			return ctx.Status(
				fiberErr.Code,
			).JSON(
				model.WebResponse[any]{
					Message: fiberErr.Message,
				},
			)
		}

		// =====================================================
		// UNHANDLED ERROR
		// =====================================================

		logrus.WithError(
			err,
		).WithFields(
			logrus.Fields{
				"path":   ctx.Path(),
				"method": ctx.Method(),
				"ip":     ctx.IP(),
			},
		).Error(
			"unhandled error",
		)

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
