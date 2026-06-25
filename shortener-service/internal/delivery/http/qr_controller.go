package http

import (
	"shortener-service/internal/model"
	"shortener-service/internal/usecase/qr"

	"github.com/gofiber/fiber/v2"
)

type QRController struct {
	useCase *qr.QRUseCase
}

func NewQRController(
	useCase *qr.QRUseCase,
) *QRController {

	return &QRController{
		useCase: useCase,
	}
}

// =====================================================
// GENERATE QR CODE
// =====================================================

func (c *QRController) Generate(
	ctx *fiber.Ctx,
) error {

	request := &model.GetQRCodeRequest{
		ID: ctx.Params(
			"id",
		),
	}

	png, err := c.useCase.Generate(
		ctx.UserContext(),
		request,
	)

	if err != nil {

		return err
	}

	ctx.Set(
		"Content-Type",
		"image/png",
	)

	ctx.Set(
		"Content-Disposition",
		"inline; filename=qrcode.png",
	)

	return ctx.Send(
		png,
	)
}
