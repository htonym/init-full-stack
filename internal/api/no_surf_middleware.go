package api

import (
	"net/http"

	"github.com/justinas/nosurf"
)

func NoSurf(inProduction bool) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		csrfHandler := nosurf.New(next)
		csrfHandler.SetBaseCookie(http.Cookie{
			HttpOnly: true,
			Path:     "/",
			Secure:   inProduction,
			SameSite: http.SameSiteLaxMode,
		})
		return csrfHandler
	}
}
