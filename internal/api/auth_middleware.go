package api

import (
	"context"
	"net/http"

	"github.com/thofftech/init-full-stack/internal/auth"
)

type ctxKey string

const userKey ctxKey = "user"

func (repo *HandlerRepo) verifyAccessToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessTokenCookie, err := r.Cookie("access-token")
		if err != nil {
			repo.homePage(w, r)
			return
		}

		idTokenCookie, err := r.Cookie("id-token")
		if err != nil {
			repo.homePage(w, r)
			return
		}

		user, err := auth.ExtractUser(idTokenCookie.Value, accessTokenCookie.Value, repo.jwksCache)
		if err != nil {
			repo.homePage(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), userKey, user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
