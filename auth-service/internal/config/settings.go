package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"time"

	"github.com/spf13/viper"
)

type Settings struct {
	Web      WebSettings
	Log      LogSettings
	Database DatabaseSettings
	Redis    RedisSettings
	SMTP     SMTPSettings

	JWT       JWTSettings
	Cache     CacheSettings
	Google    GoogleSettings
	RateLimit RateLimitSettings
}

type WebSettings struct {
	AppName string
	Port    int
	Prefork bool

	CORS CORSSettings
}

type CORSSettings struct {
	AllowOrigins     string
	AllowMethods     string
	AllowHeaders     string
	AllowCredentials bool
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

type JWTSettings struct {
	Secret string

	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
}

type CacheSettings struct {
	SessionTTL                   time.Duration
	EmailVerificationTTL         time.Duration
	EmailVerificationCooldownTTL time.Duration
	OAuthStateTTL                time.Duration
}

type GoogleSettings struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
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

type SMTPSettings struct {
	Host        string
	Port        int
	Username    string
	Password    string
	SenderName  string
	SenderEmail string
}

type LoginRateLimit struct {
	MaxAttempts int
	Window      time.Duration
}

type RegisterRateLimit struct {
	MaxAttempts int
	Window      time.Duration
}

type RateLimitSettings struct {
	Login    LoginRateLimit
	Register RegisterRateLimit
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

			CORS: CORSSettings{
				AllowOrigins: requiredString(
					v,
					"web.cors.allow_origins",
				),

				AllowMethods: requiredString(
					v,
					"web.cors.allow_methods",
				),

				AllowHeaders: requiredString(
					v,
					"web.cors.allow_headers",
				),

				AllowCredentials: v.GetBool(
					"web.cors.allow_credentials",
				),
			},
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

		JWT: JWTSettings{
			Secret: requiredString(
				v,
				"jwt.secret",
			),

			AccessTokenTTL: durationFromSeconds(
				v,
				"jwt.access_token_expired",
			),

			RefreshTokenTTL: durationFromSeconds(
				v,
				"jwt.refresh_token_expired",
			),
		},

		Cache: CacheSettings{
			SessionTTL: durationFromSeconds(
				v,
				"cache.session_ttl",
			),

			EmailVerificationTTL: durationFromSeconds(
				v,
				"cache.email_verification_ttl",
			),

			EmailVerificationCooldownTTL: durationFromSeconds(
				v,
				"cache.email_verification_cooldown",
			),

			OAuthStateTTL: durationFromSeconds(
				v,
				"cache.oauth_state_ttl",
			),
		},

		Google: GoogleSettings{
			ClientID: requiredString(
				v,
				"google.client_id",
			),

			ClientSecret: requiredString(
				v,
				"google.client_secret",
			),

			RedirectURL: requiredString(
				v,
				"google.redirect_url",
			),
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

		SMTP: SMTPSettings{

			Host: requiredString(
				v,
				"smtp.host",
			),

			Port: requiredInt(
				v,
				"smtp.port",
			),

			Username: requiredString(
				v,
				"smtp.username",
			),

			Password: requiredString(
				v,
				"smtp.password",
			),

			SenderName: requiredString(
				v,
				"smtp.sender_name",
			),

			SenderEmail: requiredString(
				v,
				"smtp.sender_email",
			),
		},

		RateLimit: RateLimitSettings{

			Login: LoginRateLimit{

				MaxAttempts: requiredInt(
					v,
					"rate_limit.login.max_attempts",
				),

				Window: durationFromSeconds(
					v,
					"rate_limit.login.window",
				),
			},

			Register: RegisterRateLimit{

				MaxAttempts: requiredInt(
					v,
					"rate_limit.register.max_attempts",
				),

				Window: durationFromSeconds(
					v,
					"rate_limit.register.window",
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
