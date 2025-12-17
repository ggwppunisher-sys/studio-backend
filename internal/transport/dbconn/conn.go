package dbconn

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/XSAM/otelsql"
	_ "github.com/jackc/pgx/v5/stdlib"
	semconv "go.opentelemetry.io/otel/semconv/v1.32.0"

	"time"
)

// SQLOptions defines configuration options for SQL database connections
type SQLOptions func(*sql.DB)

func OpenSQLConnection(ctx context.Context, dsn string, opts ...SQLOptions) (*sql.DB, error) {
	if dsn == "" {
		return nil, fmt.Errorf("dsn is required")
	}

	dbPool, err := otelsql.Open("pgx", dsn, otelsql.WithAttributes(
		semconv.DBSystemNamePostgreSQL,
	))
	if err != nil {
		return nil, fmt.Errorf("failed to open db connection: %w", err)
	}

	for _, opt := range opts {
		opt(dbPool)
	}

	if _, err = otelsql.RegisterDBStatsMetrics(
		dbPool,
		otelsql.WithAttributes(semconv.DBSystemNamePostgreSQL),
	); err != nil {
		primaryErr := fmt.Errorf("failed to register db stats metrics: %w", err)
		if closeErr := dbPool.Close(); closeErr != nil {
			return nil, errors.Join(primaryErr, fmt.Errorf("failed to close db connection: %w", closeErr))
		}
		return nil, primaryErr
	}

	if err := dbPool.PingContext(ctx); err != nil {
		primaryErr := fmt.Errorf("failed to ping db: %w", err)
		if closeErr := dbPool.Close(); closeErr != nil {
			return nil, errors.Join(primaryErr, fmt.Errorf("failed to close db connection: %w", closeErr))
		}
		return nil, primaryErr
	}

	return dbPool, nil
}

func WithMaxOpenConnections(c int) SQLOptions {
	return func(db *sql.DB) {
		if c > 0 {
			db.SetMaxOpenConns(c)
		}
	}
}

func WithMaxIdleConnections(c int) SQLOptions {
	return func(db *sql.DB) {
		if c > 0 {
			db.SetMaxIdleConns(c)
		}
	}
}

func WithConnMaxLifetime(d time.Duration) SQLOptions {
	return func(db *sql.DB) {
		if d > 0 {
			db.SetConnMaxLifetime(d)
		}
	}
}

func WithConnMaxIdleTime(d time.Duration) SQLOptions {
	return func(db *sql.DB) {
		if d > 0 {
			db.SetConnMaxIdleTime(d)
		}
	}
}

func WithConnectionTimeout(d time.Duration) SQLOptions {
	return func(db *sql.DB) {
		if d > 0 {
			db.SetConnMaxLifetime(d)
		}
	}
}
