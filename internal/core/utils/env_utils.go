package utils

import "os"

// GetEnvOrDefault returns the value of an environment variable or a default value if the environment variable is not set
func GetEnvOrDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func GetBoolEnvOrDefault(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return (value == "true" || value == "1" || value == "yes" || value == "TRUE" || value == "YES")
}

// EnvExists returns true if an environment variable exists
func EnvExists(key string) bool {
	_, ok := os.LookupEnv(key)
	return ok
}

func IsDev() bool {

	appEnv := GetEnvOrDefault("APP_ENV", "")
	isLocalhost := EnvExists("LOCALHOST")

	return (appEnv == "dev" || appEnv == "development" || isLocalhost)
}

func IsProduction() bool {
	appEnv := GetEnvOrDefault("APP_ENV", "")
	return appEnv == "production" || appEnv == "prod" || !IsDev()
}
