package config

import "github.com/spf13/viper"

func bindEnvs(v *viper.Viper) {

	// =====================================================
	// WEB
	// =====================================================

	_ = v.BindEnv(
		"web.cors.allow_origins",
		"WEB_CORS_ALLOW_ORIGINS",
	)

	// =====================================================
	// DATABASE MASTER
	// =====================================================

	_ = v.BindEnv(
		"database.master.host",
		"DB_MASTER_HOST",
	)

	_ = v.BindEnv(
		"database.master.port",
		"DB_MASTER_PORT",
	)

	_ = v.BindEnv(
		"database.master.username",
		"DB_MASTER_USERNAME",
	)

	_ = v.BindEnv(
		"database.master.password",
		"DB_MASTER_PASSWORD",
	)

	_ = v.BindEnv(
		"database.master.name",
		"DB_MASTER_NAME",
	)

	_ = v.BindEnv(
		"database.master.sslmode",
		"DB_MASTER_SSLMODE",
	)

	_ = v.BindEnv(
		"database.master.pool.idle",
		"DB_MASTER_POOL_IDLE",
	)

	_ = v.BindEnv(
		"database.master.pool.max",
		"DB_MASTER_POOL_MAX",
	)

	_ = v.BindEnv(
		"database.master.pool.lifetime",
		"DB_MASTER_POOL_LIFETIME",
	)

	// =====================================================
	// DATABASE SLAVE
	// =====================================================

	_ = v.BindEnv(
		"database.slave.host",
		"DB_SLAVE_HOST",
	)

	_ = v.BindEnv(
		"database.slave.port",
		"DB_SLAVE_PORT",
	)

	_ = v.BindEnv(
		"database.slave.username",
		"DB_SLAVE_USERNAME",
	)

	_ = v.BindEnv(
		"database.slave.password",
		"DB_SLAVE_PASSWORD",
	)

	_ = v.BindEnv(
		"database.slave.name",
		"DB_SLAVE_NAME",
	)

	_ = v.BindEnv(
		"database.slave.sslmode",
		"DB_SLAVE_SSLMODE",
	)

	_ = v.BindEnv(
		"database.slave.pool.idle",
		"DB_SLAVE_POOL_IDLE",
	)

	_ = v.BindEnv(
		"database.slave.pool.max",
		"DB_SLAVE_POOL_MAX",
	)

	_ = v.BindEnv(
		"database.slave.pool.lifetime",
		"DB_SLAVE_POOL_LIFETIME",
	)

	// =====================================================
	// REDIS
	// =====================================================

	_ = v.BindEnv(
		"redis.host",
		"REDIS_HOST",
	)

	_ = v.BindEnv(
		"redis.port",
		"REDIS_PORT",
	)

	_ = v.BindEnv(
		"redis.password",
		"REDIS_PASSWORD",
	)

	_ = v.BindEnv(
		"redis.db",
		"REDIS_DB",
	)

	// =====================================================
	// SMTP
	// =====================================================

	_ = v.BindEnv(
		"smtp.host",
		"SMTP_HOST",
	)

	_ = v.BindEnv(
		"smtp.port",
		"SMTP_PORT",
	)

	_ = v.BindEnv(
		"smtp.username",
		"SMTP_USERNAME",
	)

	_ = v.BindEnv(
		"smtp.password",
		"SMTP_PASSWORD",
	)

	// =====================================================
	// JWT
	// =====================================================

	_ = v.BindEnv(
		"jwt.secret",
		"JWT_SECRET",
	)

	// =====================================================
	// GOOGLE
	// =====================================================

	_ = v.BindEnv(
		"google.client_id",
		"GOOGLE_CLIENT_ID",
	)

	_ = v.BindEnv(
		"google.client_secret",
		"GOOGLE_CLIENT_SECRET",
	)

	_ = v.BindEnv(
		"google.redirect_url",
		"GOOGLE_REDIRECT_URL",
	)
}
