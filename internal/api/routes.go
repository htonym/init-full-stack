package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Route("/api", func(r chi.Router) {
		r.Get("/status", statusRoute)
	})

	return router
}

func statusRoute(w http.ResponseWriter, r *http.Request) {
	payload := map[string]string{
		"status":  "OK",
		"version": "1.0.0",
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
