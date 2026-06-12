package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Settings struct {
	JWTSecret string

	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration

	SessionTTL time.Duration

	LoginMaxAttempts int
	LoginWindowTTL   time.Duration

	RegisterMaxAttempts int
	RegisterWindowTTL   time.Duration
}

func NewSettings(
	v *viper.Viper,
) *Settings {

	settings := &Settings{
		JWTSecret: requiredString(
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

		SessionTTL: durationFromSeconds(
			v,
			"cache.session_ttl",
		),

		LoginMaxAttempts: requiredInt(
			v,
			"rate_limit.login.max_attempts",
		),

		LoginWindowTTL: durationFromSeconds(
			v,
			"rate_limit.login.window",
		),

		RegisterMaxAttempts: requiredInt(
			v,
			"rate_limit.register.max_attempts",
		),

		RegisterWindowTTL: durationFromSeconds(
			v,
			"rate_limit.register.window",
		),
	}

	return settings
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
