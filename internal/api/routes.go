package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/thofftech/init-full-stack/internal/auth"
	"github.com/thofftech/init-full-stack/internal/config"
	"github.com/thofftech/init-full-stack/internal/repository"
	"github.com/thofftech/init-full-stack/internal/repository/dbrepo"
)

type HandlerRepo struct {
	cfg       *config.AppConfig
	auth      *auth.Authenticator
	jwksCache *auth.JWKSCache
	DB        repository.DatabaseRepo
}

func NewRouter(cfg *config.AppConfig, authenticator *auth.Authenticator, jwksCache *auth.JWKSCache) *chi.Mux {
	repo := HandlerRepo{
		cfg:       cfg,
		auth:      authenticator,
		jwksCache: jwksCache,
		DB:        dbrepo.NewPostgresRepo(cfg.DBPool, cfg),
	}

	router := chi.NewRouter()

	router.Use(middleware.Recoverer)
	router.Use(NoSurf(cfg.Environment != config.EnvLocal))

	// Use verbose logging middleware only when running locally.
	if cfg.Environment == config.EnvLocal {
		router.Use(middleware.Logger)
	}

	fileServer := http.FileServer(http.Dir("./web/static"))
	router.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	router.With(repo.verifyAccessToken).Group(func(r chi.Router) {
		r.Get("/", repo.homePage)
		r.Get("/widgets", repo.listWidgetsPage)
		r.Get("/profile", repo.myProfilePage)
	})

	router.Route("/api", func(r chi.Router) {
		r.Get("/status", repo.appStatus)
		r.Get("/login", repo.loginHandler)
		r.Get("/callback", repo.callbackHandler)
		r.Get("/logout", repo.logoutHandler)
	})

	router.NotFound(repo.NotFoundPage)

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
