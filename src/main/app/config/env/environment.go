package env

import (
	"os"
	"strings"
)

type Env int

const (
	DEV Env = iota
	PROD
)

func (e Env) String() string {
	return []string{
		"dev",
		"prod",
	}[e]
}

// GetScope
// Get scope variable from System. Example for test.pets-api.internal.com is test.
func GetScope() string {
	return strings.ToLower(os.Getenv("SCOPE"))
}

// GetEnv
// * Get environment name from System. Priority order is as follows:
// * 1. It looks in "app.env" system property.
// * 2. If empty, it looks in SCOPE system env variable
// *		2.1 If empty, it is a local environment
// *		2.2 If not empty and starts with "test", it is a test environment
// *		2.3 Otherwise, it is a "prod" environment.
func GetEnv() string {
	env := os.Getenv("app.env")
	if !IsEmpty(env) {
		return env
	}
	env = os.Getenv("app_env")
	if !IsEmpty(env) {
		return env
	}
	scope := GetScope()
	if IsEmpty(scope) {
		return DEV.String()
	}
	return PROD.String()
}

func IsDev() bool {
	return DEV.String() == GetEnv()
}

func IsEmpty(value string) bool {
	return value == ""
}
