package config

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func NewDB(cfg *PostgresConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open(cfg.Driver, cfg.ConnectionURL())
	if err != nil {
		return nil, errors.Wrapf(err, "[postgres.NewDB] failed to initialize postgres")
	}

	if err = db.Ping(); err != nil {
		return nil, errors.Wrap(err, "[postgres.NewDB] failed to ping postgres")
	}

	db.SetMaxIdleConns(cfg.MaxIdleConnections)
	db.SetMaxOpenConns(cfg.MaxOpenConnections)

	return db, nil
}

func (cfg *PostgresConfig) ConnectionURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DatabaseName)
}
