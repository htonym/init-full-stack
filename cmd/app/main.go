package main

import (
	"context"
	"log"
	"net/http"

	"github.com/thofftech/init-full-stack/internal/api"
	"github.com/thofftech/init-full-stack/internal/auth"
	"github.com/thofftech/init-full-stack/internal/config"
)

func main() {
	ctx := context.Background()

	cfg, err := config.NewAppConfig(ctx)
	if err != nil {
		log.Fatalf("Config initialization failed before starting server: %v", err)
	}

	log.Print(cfg)

	authenticator, err := auth.NewAuthenticator(ctx, cfg)
	if err != nil {
		log.Fatalf("Failed to create authenticator instance: %v", err)
	}

	jwksCache := auth.NewJWKSCache(cfg.OAuth.JwksURL)
	if err = jwksCache.RefreshJWKS(); err != nil {
		log.Fatalf("Failed to create jwksCache instance: %v", err)
	}

	router := api.NewRouter(cfg, authenticator, jwksCache)

	log.Println("Running server...")
	err = http.ListenAndServe(":"+cfg.Port, router)
	if err != nil {
		panic(err)
	}
}
