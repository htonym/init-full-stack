package api

import (
	"context"
	"net/http"
)

type ctxKey string

const tokenKey ctxKey = "token"

func (repo *HandlerRepo) verifyAccessToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessTokenCookie, err := r.Cookie("access-token")
		if err != nil {
			repo.homePage(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), tokenKey, accessTokenCookie.Value)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
