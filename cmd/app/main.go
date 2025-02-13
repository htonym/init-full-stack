package main

import (
	"log"
	"net/http"

	"github.com/thofftech/init-full-stack/internal/api"
	"github.com/thofftech/init-full-stack/internal/config"
)

func main() {
	cfg, err := config.NewAppConfig()
	if err != nil {
		log.Fatalf("Config initialization failed before starting server: %v", err)
	}

	log.Print(cfg)

	router := api.NewRouter(cfg)

	log.Println("Running server...")
	err = http.ListenAndServe(":"+cfg.Port, router)
	if err != nil {
		panic(err)
	}
}
