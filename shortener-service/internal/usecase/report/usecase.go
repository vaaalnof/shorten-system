package report

import (
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"

	"shortener-service/internal/usecase/port"
)

type ReportUseCase struct {
	validate   *validator.Validate
	urlRepo    port.URLRepository
	reportRepo port.ReportRepository
	log        *logrus.Logger
}

func NewReportUseCase(
	validate *validator.Validate,
	urlRepo port.URLRepository,
	reportRepo port.ReportRepository,
	log *logrus.Logger,
) *ReportUseCase {

	return &ReportUseCase{
		validate:   validate,
		urlRepo:    urlRepo,
		reportRepo: reportRepo,
		log:        log,
	}
}
