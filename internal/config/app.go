package config

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Version     string
	Environment Environment
	Port        string

	AWSProfile string
	AWSRegion  string

	OAuthClientID          string
	OAuthClientSecret      string
	OAuthDomain            string
	OAuthScope             string
	OAuthCallbackURL       string
	OAuthAudience          string
	OAuthIssuerURL         string
	OAuthLogoutRedirectURL string
	OAuthJwksURL           string
}

type AWSConfig struct {
	Profile   string
	AWSRegion string
}

func NewAppConfig() (*AppConfig, error) {
	godotenv.Load()

	rawEnvValue := os.Getenv("APP_ENVIRONMENT")
	env := getEnv(rawEnvValue)
	if env == EnvUnknown {
		return nil, fmt.Errorf("unknown environment: APP_ENVIRONMENT is set to %q", rawEnvValue)
	}

	cfg := &AppConfig{
		Environment: env,
		Version:     os.Getenv("APP_VERSION"),
		Port:        os.Getenv("APP_PORT"),
		AWSProfile:  os.Getenv("APP_AWS_PROFILE"),
		AWSRegion:   os.Getenv("APP_AWS_REGION"),

		OAuthClientID:          os.Getenv("APP_OAUTH_CLIENT_ID"),
		OAuthClientSecret:      os.Getenv("APP_OAUTH_CLIENT_SECRET"),
		OAuthDomain:            os.Getenv("APP_OAUTH_DOMAIN"),
		OAuthScope:             os.Getenv("APP_OAUTH_SCOPE"),
		OAuthCallbackURL:       os.Getenv("APP_OAUTH_CALLBACK_URL"),
		OAuthAudience:          os.Getenv("APP_OAUTH_AUDIENCE"),
		OAuthIssuerURL:         os.Getenv("APP_OAUTH_ISSUER_URL"),
		OAuthLogoutRedirectURL: os.Getenv("APP_OAUTH_LOGOUT_REDIRECT_URL"),
		OAuthJwksURL:           os.Getenv("APP_OAUTH_JWKS_URL"),
	}

	// Only read version from file when running locally
	if cfg.Version == "" && cfg.Environment == EnvLocal {
		versionBytes, err := os.ReadFile("VERSION")
		if err != nil {
			return nil, fmt.Errorf("failed to read version file %w", err)
		} else {
			cfg.Version = string(versionBytes)
		}
	}

	return cfg, nil
}

func (cfg AppConfig) String() string {
	var result strings.Builder
	v := reflect.ValueOf(cfg)
	t := v.Type()

	result.WriteString("App Config:\n")
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()

		result.WriteString(fmt.Sprintf("  %s: %v\n", field.Name, value))
	}

	return result.String()
}
