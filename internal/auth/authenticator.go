package auth

import (
	"context"
	"fmt"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/thofftech/init-full-stack/internal/config"
	"golang.org/x/oauth2"
)

type Authenticator struct {
	*oidc.Provider
	oauth2.Config
}

func NewAuthenticator(cfg *config.AppConfig) (*Authenticator, error) {
	provider, err := oidc.NewProvider(context.TODO(), cfg.OAuthIssuerURL)
	if err != nil {
		return nil, fmt.Errorf("creating oidc provider for initial setup: %w", err)
	}

	conf := oauth2.Config{
		ClientID:     cfg.OAuthClientID,
		ClientSecret: cfg.OAuthClientSecret,
		RedirectURL:  cfg.OAuthCallbackURL,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID}, // Specifying only openid returns all openid info
	}

	return &Authenticator{
		Provider: provider,
		Config:   conf,
	}, nil
}
