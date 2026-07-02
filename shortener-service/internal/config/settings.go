package config

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Settings struct {
	Web       WebSettings
	Log       LogSettings
	Database  DatabaseSettings
	Redis     RedisSettings
	Cache     CacheSettings
	GeoIP     GeoIPSettings
	NATS      NATSSettings
	Auth      AuthServiceSettings
	Shortener ShortenerSettings
}

type WebSettings struct {
	AppName string
	Port    int
	Prefork bool
}

type LogSettings struct {
	Level logrus.Level
}

type DatabaseSettings struct {
	Master PostgresSettings
	Slave  PostgresSettings
}

type PostgresSettings struct {
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	SSLMode  string

	Pool DatabasePoolSettings
}

type DatabasePoolSettings struct {
	Idle     int
	Max      int
	Lifetime time.Duration
}

type RedisSettings struct {
	Host     string
	Port     string
	Password string
	DB       int

	PoolSize     int
	MinIdleConns int

	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type CacheSettings struct {
	UrlTTL time.Duration
}

type GeoIPSettings struct {
	DatabasePath string
}

type AuthServiceSettings struct {
	BaseURL string
	Timeout time.Duration
}

type ShortenerSettings struct {
	BaseURL string
}

type NATSSettings struct {
	URL string

	Timeout           time.Duration
	ReconnectWait     time.Duration
	MaxReconnects     int
	ReconnectBufferMB int

	Analytics NATSAnalyticsSettings
}

type NATSAnalyticsSettings struct {
	Replicas     int
	MaxDeliver   int
	AckWait      time.Duration
	MaxAge       time.Duration
	FetchBatch   int
	FetchTimeout time.Duration
}

func NewSettings(v *viper.Viper) *Settings {

	return &Settings{

		Web: WebSettings{

			AppName: requiredString(
				v,
				"app.name",
			),

			Port: requiredInt(
				v,
				"web.port",
			),

			Prefork: v.GetBool(
				"web.prefork",
			),
		},

		Log: LogSettings{

			Level: logrus.Level(
				requiredInt(
					v,
					"log.level",
				),
			),
		},

		Database: DatabaseSettings{

			Master: PostgresSettings{

				Host: requiredString(
					v,
					"database.master.host",
				),

				Port: requiredInt(
					v,
					"database.master.port",
				),

				Username: requiredString(
					v,
					"database.master.username",
				),

				Password: requiredString(
					v,
					"database.master.password",
				),

				Name: requiredString(
					v,
					"database.master.name",
				),

				SSLMode: requiredString(
					v,
					"database.master.sslmode",
				),

				Pool: DatabasePoolSettings{

					Idle: requiredInt(
						v,
						"database.master.pool.idle",
					),

					Max: requiredInt(
						v,
						"database.master.pool.max",
					),

					Lifetime: durationFromSeconds(
						v,
						"database.master.pool.lifetime",
					),
				},
			},

			Slave: PostgresSettings{

				Host: requiredString(
					v,
					"database.slave.host",
				),

				Port: requiredInt(
					v,
					"database.slave.port",
				),

				Username: requiredString(
					v,
					"database.slave.username",
				),

				Password: requiredString(
					v,
					"database.slave.password",
				),

				Name: requiredString(
					v,
					"database.slave.name",
				),

				SSLMode: requiredString(
					v,
					"database.slave.sslmode",
				),

				Pool: DatabasePoolSettings{

					Idle: requiredInt(
						v,
						"database.slave.pool.idle",
					),

					Max: requiredInt(
						v,
						"database.slave.pool.max",
					),

					Lifetime: durationFromSeconds(
						v,
						"database.slave.pool.lifetime",
					),
				},
			},
		},

		Redis: RedisSettings{

			Host: requiredString(
				v,
				"redis.host",
			),

			Port: requiredString(
				v,
				"redis.port",
			),

			Password: v.GetString(
				"redis.password",
			),

			DB: v.GetInt(
				"redis.db",
			),

			PoolSize:     10,
			MinIdleConns: 2,

			ReadTimeout:  5 * time.Second,
			WriteTimeout: 5 * time.Second,
		},

		Cache: CacheSettings{

			UrlTTL: durationFromSeconds(
				v,
				"cache.url_ttl_seconds",
			),
		},

		GeoIP: GeoIPSettings{

			DatabasePath: requiredString(
				v,
				"geoip.database_path",
			),
		},

		Auth: AuthServiceSettings{

			BaseURL: requiredString(
				v,
				"auth_service.base_url",
			),

			Timeout: durationFromSeconds(
				v,
				"auth_service.timeout_seconds",
			),
		},

		Shortener: ShortenerSettings{

			BaseURL: requiredString(
				v,
				"shortener.base_url",
			),
		},

		NATS: NATSSettings{

			URL: requiredString(
				v,
				"nats.url",
			),

			Timeout: durationFromSeconds(
				v,
				"nats.timeout_seconds",
			),

			ReconnectWait: durationFromSeconds(
				v,
				"nats.reconnect_wait_seconds",
			),

			MaxReconnects: v.GetInt(
				"nats.max_reconnects",
			),

			ReconnectBufferMB: requiredInt(
				v,
				"nats.reconnect_buffer_mb",
			),

			Analytics: NATSAnalyticsSettings{

				Replicas: requiredInt(
					v,
					"nats.analytics.replicas",
				),

				MaxDeliver: requiredInt(
					v,
					"nats.analytics.max_deliver",
				),

				AckWait: durationFromSeconds(
					v,
					"nats.analytics.ack_wait_seconds",
				),

				MaxAge: durationFromHours(
					v,
					"nats.analytics.max_age_hours",
				),

				FetchBatch: requiredInt(
					v,
					"nats.analytics.fetch_batch",
				),

				FetchTimeout: durationFromSeconds(
					v,
					"nats.analytics.fetch_timeout_seconds",
				),
			},
		},
	}
}

func durationFromSeconds(
	v *viper.Viper,
	key string,
) time.Duration {

	value := v.GetInt(key)

	if value <= 0 {
		panic(
			fmt.Sprintf(
				"invalid config: %s",
				key,
			),
		)
	}

	return time.Duration(value) * time.Second
}

func durationFromHours(
	v *viper.Viper,
	key string,
) time.Duration {

	value := v.GetInt(key)

	if value <= 0 {
		panic(
			fmt.Sprintf(
				"invalid config: %s",
				key,
			),
		)
	}

	return time.Duration(value) * time.Hour
}

func requiredString(
	v *viper.Viper,
	key string,
) string {

	value := v.GetString(key)

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

	value := v.GetInt(key)

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
