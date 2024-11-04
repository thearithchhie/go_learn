package utils

import (
	"os"
	"strconv"
)

// GetEnv retrieves the value of the environment variable named by the key.
// If the variable is not set, it returns the provided default value.
func GetEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return ""
}

func GetenvInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}
