package config

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Version     string
	Environment Environment
	Port        string

	RemoteParamPath string

	AWSProfile string
	AWSRegion  string

	OAuth OauthConfig
}

type OauthConfig struct {
	ClientID          string `json:"APP_OAUTH_CLIENT_ID"`
	ClientSecret      string `json:"APP_OAUTH_CLIENT_SECRET"`
	Domain            string `json:"APP_OAUTH_DOMAIN"`
	Scope             string `json:"APP_OAUTH_SCOPE"`
	CallbackURL       string `json:"APP_OAUTH_CALLBACK_URL"`
	IssuerURL         string `json:"APP_OAUTH_ISSUER_URL"`
	Audience          string `json:"APP_OAUTH_AUDIENCE"`
	LogoutRedirectURL string `json:"APP_OAUTH_LOGOUT_REDIRECT_URL"`
	JwksURL           string `json:"APP_OAUTH_JWKS_URL"`
}

type AWSConfig struct {
	Profile   string
	AWSRegion string
}

func NewAppConfig(ctx context.Context) (*AppConfig, error) {
	godotenv.Load()

	rawEnvValue := os.Getenv("APP_ENVIRONMENT")
	env := getEnv(rawEnvValue)
	if env == EnvUnknown {
		return nil, fmt.Errorf("unknown environment: APP_ENVIRONMENT is set to %q", rawEnvValue)
	}

	cfg := &AppConfig{
		Environment:     env,
		Version:         os.Getenv("APP_VERSION"),
		Port:            os.Getenv("APP_PORT"),
		AWSProfile:      os.Getenv("APP_AWS_PROFILE"),
		AWSRegion:       os.Getenv("APP_AWS_REGION"),
		RemoteParamPath: os.Getenv("APP_REMOTE_PARAM_PATH"),
	}

	if cfg.Environment == EnvLocal {
		if err := cfg.loadLocalConfig(); err != nil {
			return nil, err
		}
	} else {
		if err := cfg.fetchRemoteConfig(ctx); err != nil {
			return nil, err
		}
	}

	return cfg, nil
}

func (cfg *AppConfig) fetchRemoteConfig(ctx context.Context) error {
	awsCfg, err := awsconfig.LoadDefaultConfig(ctx,
		awsconfig.WithRegion(cfg.AWSRegion),
	)
	if err != nil {
		return fmt.Errorf("unable to load SDK config: %w", err)
	}

	client := ssm.NewFromConfig(awsCfg)

	fmt.Println("cfg.RemoteParamPath:", cfg.RemoteParamPath)

	out, err := client.GetParameter(ctx, &ssm.GetParameterInput{
		Name:           &cfg.RemoteParamPath,
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		return fmt.Errorf("unable to get parameter: %w", err)
	}

	if out.Parameter == nil || out.Parameter.Value == nil {
		return fmt.Errorf("parameter not found or has no value")
	}

	if err := json.Unmarshal([]byte(*out.Parameter.Value), &cfg.OAuth); err != nil {
		panic(fmt.Errorf("failed to unmarshal JSON: %w", err))
	}

	return nil
}

func (cfg *AppConfig) loadLocalConfig() error {
	cfg.OAuth.ClientID = os.Getenv("APP_OAUTH_CLIENT_ID")
	cfg.OAuth.ClientSecret = os.Getenv("APP_OAUTH_CLIENT_SECRET")
	cfg.OAuth.Domain = os.Getenv("APP_OAUTH_DOMAIN")
	cfg.OAuth.Scope = os.Getenv("APP_OAUTH_SCOPE")
	cfg.OAuth.CallbackURL = os.Getenv("APP_OAUTH_CALLBACK_URL")
	cfg.OAuth.IssuerURL = os.Getenv("APP_OAUTH_ISSUER_URL")
	cfg.OAuth.Audience = os.Getenv("APP_OAUTH_AUDIENCE")
	cfg.OAuth.LogoutRedirectURL = os.Getenv("APP_OAUTH_LOGOUT_REDIRECT_URL")
	cfg.OAuth.JwksURL = os.Getenv("APP_OAUTH_JWKS_URL")

	versionBytes, err := os.ReadFile("VERSION")
	if err != nil {
		return fmt.Errorf("failed to read version file %w", err)
	} else {
		cfg.Version = string(versionBytes)
	}

	return nil
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

func (cfg OauthConfig) String() string {
	var result strings.Builder
	v := reflect.ValueOf(cfg)
	t := v.Type()

	result.WriteString("\n")
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()

		if strings.Contains(strings.ToLower(field.Name), "secret") {
			valueStr, ok := value.(string)
			if ok {
				value = valueStr[0:3] + "******"
			}
		}

		result.WriteString(fmt.Sprintf("    %s: %v\n", field.Name, value))
	}

	return result.String()
}
