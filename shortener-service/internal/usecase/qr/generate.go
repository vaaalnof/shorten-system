package qr

import (
	"bytes"
	"context"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"strings"

	"github.com/nfnt/resize"
	"github.com/skip2/go-qrcode"

	"shortener-service/internal/delivery/http/middleware"
	"shortener-service/internal/exception"
	"shortener-service/internal/model"
)

// =====================================================
// GENERATE QR CODE
// =====================================================

func (u *QRUseCase) Generate(
	ctx context.Context,
	request *model.GetQRCodeRequest,
) ([]byte, error) {

	if err := u.validate.Struct(
		request,
	); err != nil {

		return nil, exception.Validation(
			err,
		)
	}

	meta := middleware.GetMeta(
		ctx,
	)

	if meta == nil ||
		meta.UserID == "" {

		return nil, exception.Unauthorized(
			"unauthorized",
		)
	}

	// =====================================================
	// URL
	// =====================================================

	url, err := u.urlRepo.FindByID(
		ctx,
		request.ID,
	)

	if err != nil {

		return nil, exception.Internal(
			"failed to generate qrcode",
		)
	}

	if url == nil {

		return nil, exception.NotFound(
			"shorturl not found",
		)
	}

	if url.UserID != meta.UserID {

		return nil, exception.NotFound(
			"shorturl not found",
		)
	}

	// =====================================================
	// SHORT URL
	// =====================================================

	shortURL := strings.TrimRight(
		u.baseURL,
		"/",
	) + "/" + url.ShortCode

	// =====================================================
	// QR
	// =====================================================

	qr, err := qrcode.New(
		shortURL,
		qrcode.High,
	)

	if err != nil {

		return nil, exception.Internal(
			"failed to generate qrcode",
		)
	}

	qr.ForegroundColor = color.RGBA{
		R: 37,
		G: 99,
		B: 235,
		A: 255,
	}

	qr.BackgroundColor = color.White

	const size = 512

	qrImage := qr.Image(
		size,
	)

	// =====================================================
	// LOGO
	// =====================================================

	logoFile, err := os.Open(
		"assets/logo.png",
	)

	if err == nil {

		defer logoFile.Close()

		logoImage, err := png.Decode(
			logoFile,
		)

		if err == nil {

			logoSize := uint(
				size / 5,
			)

			logoImage = resize.Resize(
				logoSize,
				logoSize,
				logoImage,
				resize.Lanczos3,
			)

			canvas := image.NewRGBA(
				qrImage.Bounds(),
			)

			draw.Draw(
				canvas,
				canvas.Bounds(),
				qrImage,
				image.Point{},
				draw.Src,
			)

			logoBounds := logoImage.Bounds()

			x := (canvas.Bounds().Dx() - logoBounds.Dx()) / 2

			y := (canvas.Bounds().Dy() - logoBounds.Dy()) / 2

			draw.Draw(
				canvas,
				image.Rect(
					x,
					y,
					x+logoBounds.Dx(),
					y+logoBounds.Dy(),
				),
				logoImage,
				image.Point{},
				draw.Over,
			)

			buf := new(
				bytes.Buffer,
			)

			if err := png.Encode(
				buf,
				canvas,
			); err != nil {

				return nil, exception.Internal(
					"failed to generate qrcode",
				)
			}

			return buf.Bytes(), nil
		}
	}

	// =====================================================
	// FALLBACK WITHOUT LOGO
	// =====================================================

	pngBytes, err := qr.PNG(
		size,
	)

	if err != nil {

		return nil, exception.Internal(
			"failed to generate qrcode",
		)
	}

	return pngBytes, nil
}
