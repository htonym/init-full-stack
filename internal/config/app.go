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

	result.WriteString("App Config\n")
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()

		result.WriteString(fmt.Sprintf("  %s: %v\n", field.Name, value))
	}

	return result.String()
}
