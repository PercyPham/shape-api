package config

import (
	"os"
	"strconv"
)

// Load retrieves configs from environment variables.
//
// Call this function as close as possible to the start of your program (ideally in main).
func Load() {
	loadAppConfig()
	loadMySQLConfig()
}

// getENV returns the environment variable with matching key,
// returns defaultVal if not found.
func getENV(key, defaultVal string) string {
	env := os.Getenv(key)
	if env == "" {
		return defaultVal
	}
	return env
}

// getIntENV returns the interger environment variable with matching key,
// returns defaultVal if not found or found non integer number.
func getIntENV(key string, defaultVal int) int {
	env := getENV(key, "")
	envInt, err := strconv.Atoi(env)
	if err != nil {
		return defaultVal
	}
	return envInt
}
