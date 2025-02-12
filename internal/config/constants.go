package config

import "strings"

type Environment int

const (
	EnvLocal Environment = iota
	EnvDev
	EnvTest
	EnvNonProd
	EnvStaging
	EnvProd
	EnvUnknown
)

func (e Environment) String() string {
	return [...]string{
		"LOCAL",
		"DEV",
		"TEST",
		"NON-PROD",
		"STAGING",
		"PROD",
		"UNKNOWN",
	}[e]
}

func getEnv(env string) Environment {
	switch strings.ToLower(env) {
	case "local":
		return EnvLocal
	case "dev":
		return EnvDev
	case "test":
		return EnvTest
	case "nonprod":
		return EnvNonProd
	case "non-prod":
		return EnvNonProd
	case "staging":
		return EnvStaging
	case "prod":
		return EnvProd
	default:
		return EnvUnknown
	}
}
