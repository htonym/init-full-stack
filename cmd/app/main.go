package main

import (
	"log"
	"net/http"

	"github.com/thofftech/init-full-stack/internal/api"
	"github.com/thofftech/init-full-stack/internal/auth"
	"github.com/thofftech/init-full-stack/internal/config"
)

func main() {
	cfg, err := config.NewAppConfig()
	if err != nil {
		log.Fatalf("Config initialization failed before starting server: %v", err)
	}

	log.Print(cfg)

	authenticator, err := auth.NewAuthenticator(cfg)
	if err != nil {
		log.Fatalf("Failed to create authenticator instance: %v", err)
	}

	jwksCache := auth.NewJWKSCache(cfg.OAuthJwksURL)
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
