package config

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Database struct {
	Master *sql.DB
	Slave  *sql.DB
	Log    *logrus.Logger
}

func NewDatabase(
	v *viper.Viper,
	log *logrus.Logger,
) *Database {

	master := newPostgresConnection(
		v,
		log,
		"database.master",
	)

	slave := newPostgresConnection(
		v,
		log,
		"database.slave",
	)

	return &Database{
		Master: master,
		Slave:  slave,
		Log:    log,
	}
}

func newPostgresConnection(
	v *viper.Viper,
	log *logrus.Logger,
	prefix string,
) *sql.DB {

	username := v.GetString(
		fmt.Sprintf("%s.username", prefix),
	)

	password := v.GetString(
		fmt.Sprintf("%s.password", prefix),
	)

	host := v.GetString(
		fmt.Sprintf("%s.host", prefix),
	)

	port := v.GetInt(
		fmt.Sprintf("%s.port", prefix),
	)

	database := v.GetString(
		fmt.Sprintf("%s.name", prefix),
	)

	sslmode := v.GetString(
		fmt.Sprintf("%s.sslmode", prefix),
	)

	idleConnection := v.GetInt(
		fmt.Sprintf("%s.pool.idle", prefix),
	)

	maxConnection := v.GetInt(
		fmt.Sprintf("%s.pool.max", prefix),
	)

	maxLifeTimeConnection := v.GetInt(
		fmt.Sprintf("%s.pool.lifetime", prefix),
	)

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		host,
		port,
		username,
		password,
		database,
		sslmode,
	)

	db, err := sql.Open(
		"postgres",
		dsn,
	)

	if err != nil {

		log.Fatalf(
			"failed to open PostgreSQL connection (%s): %v",
			prefix,
			err,
		)
	}

	db.SetMaxIdleConns(
		idleConnection,
	)

	db.SetMaxOpenConns(
		maxConnection,
	)

	db.SetConnMaxLifetime(
		time.Second * time.Duration(maxLifeTimeConnection),
	)

	if err := db.Ping(); err != nil {

		log.Fatalf(
			"failed to ping PostgreSQL (%s): %v",
			prefix,
			err,
		)
	}

	log.Infof(
		"connected PostgreSQL (%s) at %s:%d/%s",
		prefix,
		host,
		port,
		database,
	)

	return db
}
