package qr

import (
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"

	"shortener-service/internal/usecase/port"
)

type QRUseCase struct {
	validate *validator.Validate
	urlRepo  port.URLRepository
	baseURL  string
	log      *logrus.Logger
}

func NewQRUseCase(
	validate *validator.Validate,
	urlRepo port.URLRepository,
	baseURL string,
	log *logrus.Logger,
) *QRUseCase {

	return &QRUseCase{
		validate: validate,
		urlRepo:  urlRepo,
		baseURL:  baseURL,
		log:      log,
	}
}
