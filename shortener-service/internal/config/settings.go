package config

import (
	"time"

	"github.com/spf13/viper"
)

type Settings struct {

	// =====================================================
	// AUTH SERVICE
	// =====================================================

	AuthServiceBaseURL string

	AuthServiceTimeout time.Duration

	// =====================================================
	// CACHE
	// =====================================================

	URLCacheTTL time.Duration

	// =====================================================
	// SHORTENER
	// =====================================================
	ShortenerBaseURL string

	// =====================================================
	// GEO IP
	// =====================================================

	GeoIPDatabasePath string

	// =====================================================
	// ANALYTICS (NATS)
	// =====================================================

	NATSAnalyticsReplicas int

	NATSAnalyticsMaxDeliver int

	NATSAnalyticsAckWait time.Duration

	NATSAnalyticsMaxAge time.Duration

	NATSAnalyticsFetchBatch int

	NATSAnalyticsFetchTimeout time.Duration
}

func NewSettings(
	v *viper.Viper,
) *Settings {

	return &Settings{

		// =====================================================
		// AUTH SERVICE
		// =====================================================

		AuthServiceBaseURL: v.GetString(
			"auth_service.base_url",
		),

		AuthServiceTimeout: time.Duration(
			v.GetInt(
				"auth_service.timeout_seconds",
			),
		) * time.Second,

		// =====================================================
		// CACHE
		// =====================================================

		URLCacheTTL: time.Duration(
			v.GetInt(
				"cache.url_ttl_seconds",
			),
		) * time.Second,

		// =====================================================
		// SHORTENER
		// =====================================================
		ShortenerBaseURL: v.GetString(
			"shortener.base_url",
		),

		// =====================================================
		// GEO IP
		// =====================================================

		GeoIPDatabasePath: v.GetString(
			"geoip.database_path",
		),

		// =====================================================
		// NATS ANALYTICS
		// =====================================================

		NATSAnalyticsReplicas: v.GetInt(
			"nats.analytics.replicas",
		),

		NATSAnalyticsMaxDeliver: v.GetInt(
			"nats.analytics.max_deliver",
		),

		NATSAnalyticsAckWait: time.Duration(
			v.GetInt(
				"nats.analytics.ack_wait_seconds",
			),
		) * time.Second,

		NATSAnalyticsMaxAge: time.Duration(
			v.GetInt(
				"nats.analytics.max_age_hours",
			),
		) * time.Hour,

		NATSAnalyticsFetchBatch: v.GetInt(
			"nats.analytics.fetch_batch",
		),

		NATSAnalyticsFetchTimeout: time.Duration(
			v.GetInt(
				"nats.analytics.fetch_timeout_seconds",
			),
		) * time.Second,
	}
}
