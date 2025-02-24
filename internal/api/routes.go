package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/thofftech/init-full-stack/internal/config"
)

type HandlerRepo struct {
	cfg *config.AppConfig
}

func NewRouter(cfg *config.AppConfig) *chi.Mux {
	repo := HandlerRepo{
		cfg: cfg,
	}

	router := chi.NewRouter()

	// Use verbose logging middleware only when running locally.
	if cfg.Environment == config.EnvLocal {
		router.Use(middleware.Logger)
	}

	fileServer := http.FileServer(http.Dir("./web/static"))
	router.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	router.Get("/", homePage)

	router.Route("/api", func(r chi.Router) {
		r.Get("/status", repo.appStatus)
	})

	return router
}

func (repo *HandlerRepo) appStatus(w http.ResponseWriter, r *http.Request) {
	payload := map[string]string{
		"status":      "OK",
		"version":     repo.cfg.Version,
		"environment": repo.cfg.Environment.String(),
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
