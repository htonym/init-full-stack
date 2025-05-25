package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/thofftech/init-full-stack/internal/config"
)

func main() {
	ctx := context.Background()

	cfg, err := config.NewAppConfig(ctx)
	if err != nil {
		log.Fatalf("Config initialization failed before starting server: %v", err)
	}

	fmt.Println(cfg)

	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Remote.DbUser,
		cfg.Remote.DbPassword,
		cfg.Remote.DbHost,
		cfg.Remote.DbPort,
		cfg.Remote.DbName,
		cfg.Remote.DbSslMode,
	)

	dbpool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	var greeting string
	err = dbpool.QueryRow(context.Background(), "select 'DB Connection OK'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not connect to database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)
}
