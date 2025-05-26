package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/thofftech/init-full-stack/internal/config"
)

func main() {
	ctx := context.Background()

	cfg, err := config.NewAppConfig(ctx)
	if err != nil {
		log.Fatalf("Config initialization failed before starting server: %v", err)
	}

	fmt.Print(cfg)

	fmt.Println("Pinging database...")
	err = cfg.DBPool.Ping(ctx)
	if err != nil {
		fmt.Println("Failed to ping db")
		os.Exit(1)
	}

	fmt.Println("DB Connection OK")
}
