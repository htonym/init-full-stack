package config

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func (cfg *AppConfig) initDB(ctx context.Context) error {
	dbConfig, err := pgxpool.ParseConfig(cfg.dbConnStr())
	if err != nil {
		return fmt.Errorf("parsing pgx db pool: %w", err)
	}

	// Set pool configurations
	dbConfig.MaxConns = 20
	dbConfig.MinConns = 5
	dbConfig.MaxConnLifetime = time.Hour
	dbConfig.MaxConnIdleTime = 30 * time.Minute
	dbConfig.HealthCheckPeriod = time.Minute

	cfg.DBPool, err = pgxpool.NewWithConfig(ctx, dbConfig)
	if err != nil {
		return fmt.Errorf("attempt to create new pool: %w", err)
	}

	return nil
}

func (cfg *AppConfig) dbConnStr() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Remote.DbUser,
		cfg.Remote.DbPassword,
		cfg.Remote.DbHost,
		cfg.Remote.DbPort,
		cfg.Remote.DbName,
		cfg.Remote.DbSslMode,
	)
}
