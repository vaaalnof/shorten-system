package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func NewViper() *viper.Viper {

	// =====================================================
	// LOAD .ENV (Development Only)
	// =====================================================

	_ = godotenv.Load()
	_ = godotenv.Load("../.env")

	v := viper.New()

	// =====================================================
	// CONFIG FILE
	// =====================================================

	v.SetConfigName("config")
	v.SetConfigType("json")

	v.AddConfigPath("./")
	v.AddConfigPath("./../")

	// =====================================================
	// LOAD CONFIG FILE
	// =====================================================

	if err := v.ReadInConfig(); err != nil {

		panic(
			fmt.Errorf(
				"failed to load config: %w",
				err,
			),
		)
	}

	// =====================================================
	// ENV BINDING
	// =====================================================

	bindEnvs(
		v,
	)

	return v
}
