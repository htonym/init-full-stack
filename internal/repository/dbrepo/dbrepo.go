package dbrepo

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/thofftech/init-full-stack/internal/config"
	"github.com/thofftech/init-full-stack/internal/repository"
)

type postgresDBRepo struct {
	App  *config.AppConfig
	Pool *pgxpool.Pool
}

func NewPostgresRepo(dbPool *pgxpool.Pool, cfg *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App:  cfg,
		Pool: dbPool,
	}
}
