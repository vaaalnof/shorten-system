package shorturl

import (
	"shortener-service/internal/usecase/analytics"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"

	"shortener-service/internal/security"
	"shortener-service/internal/usecase/port"
)

type URLUseCase struct {
	validate           *validator.Validate
	urlRepo            port.URLRepository
	reservedAliasRepo  port.ReservedAliasRepository
	passwordHash       security.PasswordHash
	cache              port.Cache
	urlCacheTTL        time.Duration
	analyticsPublisher *analytics.AnalyticsEventPublisher
	log                *logrus.Logger
}

func NewURLUseCase(
	validate *validator.Validate,
	urlRepo port.URLRepository,
	reservedAliasRepo port.ReservedAliasRepository,
	passwordHash security.PasswordHash,
	cache port.Cache,
	urlCacheTTL time.Duration,
	analyticsPublisher *analytics.AnalyticsEventPublisher,
	log *logrus.Logger,
) *URLUseCase {

	return &URLUseCase{
		validate:           validate,
		urlRepo:            urlRepo,
		reservedAliasRepo:  reservedAliasRepo,
		passwordHash:       passwordHash,
		cache:              cache,
		urlCacheTTL:        urlCacheTTL,
		analyticsPublisher: analyticsPublisher,
		log:                log,
	}
}
