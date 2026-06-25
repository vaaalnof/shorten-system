package analytics

import (
	"shortener-service/internal/entity"
	"shortener-service/internal/security"
)

type AnalyticsEventEnricher struct {
	geoIP security.GeoIP
}

func NewAnalyticsEventEnricher(
	geoIP security.GeoIP,
) *AnalyticsEventEnricher {

	return &AnalyticsEventEnricher{
		geoIP: geoIP,
	}
}

func (e *AnalyticsEventEnricher) Enrich(
	event *entity.AnalyticsEvent,
) {

	if e.geoIP == nil ||
		event.IPAddress == nil {

		return
	}

	country := e.geoIP.Country(
		*event.IPAddress,
	)

	if country == "" {
		return
	}

	event.Country = &country
}
