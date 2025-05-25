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

func NewAuthenticator(ctx context.Context, cfg *config.AppConfig) (*Authenticator, error) {
	provider, err := oidc.NewProvider(ctx, cfg.Remote.IssuerURL)
	if err != nil {
		return nil, fmt.Errorf("creating oidc provider for initial setup: %w", err)
	}

	conf := oauth2.Config{
		ClientID:     cfg.Remote.ClientID,
		ClientSecret: cfg.Remote.ClientSecret,
		RedirectURL:  cfg.Remote.CallbackURL,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID}, // Specifying only openid returns all openid info
	}

	return &Authenticator{
		Provider: provider,
		Config:   conf,
	}, nil
}
