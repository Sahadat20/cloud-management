package db

import (
	"context"
	"fmt"
	"time"

	"github.com/Sahadat20/cloud-management/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Pool *pgxpool.Pool
}

func New(cfg *config.DatabaseConfig) (*DB, error) {
	poolconfig, err := pgxpool.ParseConfig(cfg.DSN())
	if err != nil {
		return nil, fmt.Errorf("failed to parse database config: %w", err)
	}

	// connection pool
	poolconfig.MaxConns = 25
	poolconfig.MinConns = 5
	poolconfig.MaxConnLifetime = time.Hour
	poolconfig.MaxConnIdleTime = 30 * time.Minute
	poolconfig.HealthCheckPeriod = time.Minute

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	pool, err := pgxpool.NewWithConfig(ctx, poolconfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	// verify connection
	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	return &DB{
		Pool: pool,
	}, nil

}

func (db *DB) Close() {
	if db.Pool != nil {
		db.Pool.Close()
	}
}

func (db *DB) Health(ctx context.Context) error {
	return db.Pool.Ping(ctx)
}
