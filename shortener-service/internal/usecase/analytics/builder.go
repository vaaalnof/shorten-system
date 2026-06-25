package analytics

import (
	"time"

	analyticsutil "shortener-service/internal/utils/analytics"

	"shortener-service/internal/delivery/http/middleware"
	"shortener-service/internal/entity"
)

func (p *AnalyticsEventPublisher) buildEvent(
	meta *middleware.RequestMeta,
	url *entity.URL,
) *entity.AnalyticsEvent {

	event := &entity.AnalyticsEvent{
		UrlID:     url.ID,
		ShortCode: url.ShortCode,
		ClickedAt: time.Now().Unix(),
	}

	if meta == nil {
		return event
	}

	p.fillReferer(
		event,
		meta,
	)

	p.fillUserAgent(
		event,
		meta,
	)

	p.fillIPAddress(
		event,
		meta,
	)

	return event
}

func (p *AnalyticsEventPublisher) fillReferer(
	event *entity.AnalyticsEvent,
	meta *middleware.RequestMeta,
) {

	if meta.Referer == "" {
		return
	}

	referer := meta.Referer

	event.Referer = &referer

	source := analyticsutil.DetectSource(
		referer,
	)

	event.Source = &source
}

func (p *AnalyticsEventPublisher) fillUserAgent(
	event *entity.AnalyticsEvent,
	meta *middleware.RequestMeta,
) {

	if meta.UserAgent == "" {
		return
	}

	userAgent := meta.UserAgent

	event.UserAgent = &userAgent

	browser := analyticsutil.DetectBrowser(
		userAgent,
	)

	event.Browser = &browser

	os := analyticsutil.DetectOS(
		userAgent,
	)

	event.OS = &os

	device := analyticsutil.DetectDevice(
		userAgent,
	)

	event.Device = &device
}

func (p *AnalyticsEventPublisher) fillIPAddress(
	event *entity.AnalyticsEvent,
	meta *middleware.RequestMeta,
) {

	if meta.IPAddress == "" {
		return
	}

	ipAddress := meta.IPAddress

	event.IPAddress = &ipAddress
}
