package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/sirupsen/logrus"
)

type Database struct {
	Master *sql.DB
	Slave  *sql.DB
	Log    *logrus.Logger
}

func NewDatabase(
	cfg DatabaseSettings,
	log *logrus.Logger,
) *Database {

	master := newPostgresConnection(
		cfg.Master,
		log,
		"master",
	)

	slave := newPostgresConnection(
		cfg.Slave,
		log,
		"slave",
	)

	return &Database{
		Master: master,
		Slave:  slave,
		Log:    log,
	}
}

func newPostgresConnection(
	cfg PostgresSettings,
	log *logrus.Logger,
	name string,
) *sql.DB {

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.Password,
		cfg.Name,
		cfg.SSLMode,
	)

	db, err := sql.Open(
		"postgres",
		dsn,
	)

	if err != nil {

		log.WithError(err).
			Fatalf(
				"failed to open PostgreSQL connection (%s)",
				name,
			)
	}

	db.SetMaxIdleConns(
		cfg.Pool.Idle,
	)

	db.SetMaxOpenConns(
		cfg.Pool.Max,
	)

	db.SetConnMaxLifetime(
		cfg.Pool.Lifetime,
	)

	if err := db.Ping(); err != nil {

		log.WithError(err).
			Fatalf(
				"failed to ping PostgreSQL (%s)",
				name,
			)
	}

	log.Infof(
		"connected PostgreSQL (%s) at %s:%d/%s",
		name,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)

	return db
}
