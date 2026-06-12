package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/sirupsen/logrus"
)

type Repository struct {
	Master *sql.DB
	Slave  *sql.DB
	Log    *logrus.Logger
}

func NewRepository(
	master *sql.DB,
	slave *sql.DB,
	log *logrus.Logger,
) *Repository {
	return &Repository{
		Master: master,
		Slave:  slave,
		Log:    log,
	}
}

func (r *Repository) QueryRow(
	ctx context.Context,
	query string, args ...interface{},
) *sql.Row {
	if r.Slave == nil {
		r.Log.Error("slave DB is not initialized")
		return &sql.Row{}
	}
	return r.Slave.QueryRowContext(ctx, query, args...)
}

func (r *Repository) Query(
	ctx context.Context,
	query string, args ...interface{},
) (*sql.Rows, error) {
	if r.Slave == nil {
		return nil, errors.New("slave DB is not initialized")
	}
	return r.Slave.QueryContext(ctx, query, args...)
}

func (r *Repository) Exec(ctx context.Context,
	query string,
	args ...interface{},
) (sql.Result, error) {
	if r.Master == nil {
		return nil, errors.New("master DB is not initialized")
	}
	return r.Master.ExecContext(ctx, query, args...)
}

func (r *Repository) QueryRowMaster(
	ctx context.Context,
	query string,
	args ...interface{},
) (*sql.Row, error) {
	if r.Master == nil {
		return nil, errors.New("master DB is not initialized")
	}
	return r.Master.QueryRowContext(ctx, query, args...), nil
}

func (r *Repository) WithTransaction(
	ctx context.Context,
	fn func(ctx context.Context, tx *sql.Tx) error,
) error {
	if r.Master == nil {
		return errors.New("master DB is not initialized")
	}

	tx, err := r.Master.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	if err := fn(ctx, tx); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
