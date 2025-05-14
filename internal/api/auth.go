package api

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"time"

	"github.com/go-chi/render"
	"github.com/golang-jwt/jwt/v5"
)

func (repo *HandlerRepo) loginHandler(w http.ResponseWriter, r *http.Request) {
	state, err := generateRandomState()
	if err != nil {
		return
	}

	newURL := repo.auth.AuthCodeURL(state)

	http.Redirect(w, r, newURL, http.StatusMovedPermanently)
}

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	// state := base64.StdEncoding.EncodeToString(b)

	return "state", nil
}

func (repo *HandlerRepo) callbackHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()

	code := r.URL.Query().Get("code")

	token, err := repo.auth.Exchange(ctx, code)
	if err != nil {
		render.JSON(w, r, map[string]string{
			"error": fmt.Sprintf("Unauthorized: Failed to exchange code for token: %v", err),
		})
		return
	}

	rawAccessToken := token.AccessToken

	_, err = repo.jwksCache.ParseToken(rawAccessToken)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			render.JSON(w, r, map[string]string{
				"error": "Unauthorized: Expired access token.",
			})
		} else {
			slog.Error(fmt.Sprintf("invalid access token: %v", err))
			render.JSON(w, r, map[string]string{
				"error": "Unauthorized: Invalid access token.",
			})
		}
		return
	}

	rawIdToken, ok := token.Extra("id_token").(string)
	if !ok {
		render.JSON(w, r, map[string]string{
			"error": "Unauthorized: Invalid access token.",
		})
	}

	setCookie(w, "access-token", rawAccessToken)
	setCookie(w, "id-token", rawIdToken)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func setCookie(w http.ResponseWriter, name, value string) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   60 * 60 * 24, // 1 Day
	}
	http.SetCookie(w, cookie)
}

func (repo *HandlerRepo) logoutHandler(w http.ResponseWriter, r *http.Request) {
	removeCookie(w, "access-token")
	removeCookie(w, "id-token")

	logoutURL, err := url.Parse(fmt.Sprintf("https://%s/logout", repo.cfg.OAuthDomain))
	if err != nil {
		render.JSON(w, r, map[string]string{
			"error": "Internal Server error:  repo.cfg.OAuthDomain parameter",
		})
	}

	params := url.Values{}
	params.Add("client_id", repo.cfg.OAuthClientID)
	params.Add("logout_uri", repo.cfg.OAuthLogoutRedirectURL)
	logoutURL.RawQuery = params.Encode()

	http.Redirect(w, r, logoutURL.String(), http.StatusTemporaryRedirect)
}

func removeCookie(w http.ResponseWriter, name string) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   -1, // Setting to -1 effectively removes the token
		Expires:  time.Unix(0, 0),
	}
	http.SetCookie(w, cookie)
}
