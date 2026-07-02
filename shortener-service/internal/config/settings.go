package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Settings struct {
	AuthServiceBaseURL string
	AuthServiceTimeout time.Duration

	URLCacheTTL time.Duration

	ShortenerBaseURL string

	GeoIPDatabasePath string

	NATSAnalyticsReplicas     int
	NATSAnalyticsMaxDeliver   int
	NATSAnalyticsAckWait      time.Duration
	NATSAnalyticsMaxAge       time.Duration
	NATSAnalyticsFetchBatch   int
	NATSAnalyticsFetchTimeout time.Duration
}

func NewSettings(
	v *viper.Viper,
) *Settings {

	return &Settings{
		AuthServiceBaseURL: requiredString(
			v,
			"auth_service.base_url",
		),

		AuthServiceTimeout: durationFromSeconds(
			v,
			"auth_service.timeout_seconds",
		),

		URLCacheTTL: durationFromSeconds(
			v,
			"cache.url_ttl_seconds",
		),

		ShortenerBaseURL: requiredString(
			v,
			"shortener.base_url",
		),

		GeoIPDatabasePath: requiredString(
			v,
			"geoip.database_path",
		),

		NATSAnalyticsReplicas: requiredInt(
			v,
			"nats.analytics.replicas",
		),

		NATSAnalyticsMaxDeliver: requiredInt(
			v,
			"nats.analytics.max_deliver",
		),

		NATSAnalyticsAckWait: durationFromSeconds(
			v,
			"nats.analytics.ack_wait_seconds",
		),

		NATSAnalyticsMaxAge: durationFromHours(
			v,
			"nats.analytics.max_age_hours",
		),

		NATSAnalyticsFetchBatch: requiredInt(
			v,
			"nats.analytics.fetch_batch",
		),

		NATSAnalyticsFetchTimeout: durationFromSeconds(
			v,
			"nats.analytics.fetch_timeout_seconds",
		),
	}
}

func durationFromSeconds(
	v *viper.Viper,
	key string,
) time.Duration {

	value := v.GetInt(
		key,
	)

	if value <= 0 {
		panic(
			fmt.Sprintf(
				"invalid config: %s",
				key,
			),
		)
	}

	return time.Duration(
		value,
	) * time.Second
}

func durationFromHours(
	v *viper.Viper,
	key string,
) time.Duration {

	value := v.GetInt(
		key,
	)

	if value <= 0 {
		panic(
			fmt.Sprintf(
				"invalid config: %s",
				key,
			),
		)
	}

	return time.Duration(
		value,
	) * time.Hour
}

func requiredString(
	v *viper.Viper,
	key string,
) string {

	value := v.GetString(
		key,
	)

	if value == "" {
		panic(
			fmt.Sprintf(
				"missing config: %s",
				key,
			),
		)
	}

	return value
}

func requiredInt(
	v *viper.Viper,
	key string,
) int {

	value := v.GetInt(
		key,
	)

	if value <= 0 {
		panic(
			fmt.Sprintf(
				"invalid config: %s",
				key,
			),
		)
	}

	return value
}
