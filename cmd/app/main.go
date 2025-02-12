package main

import (
	"log"
	"net/http"

	"github.com/thofftech/init-full-stack/internal/api"
	"github.com/thofftech/init-full-stack/internal/config"
)

func main() {
	router := api.NewRouter()

	cfg, err := config.NewAppConfig()
	if err != nil {
		log.Fatalf("Config initialization failed before starting server: %v", err)
	}

	log.Println(cfg)

	log.Println("Running server...")
	err = http.ListenAndServe(":"+cfg.Port, router)
	if err != nil {
		panic(err)
	}
}
