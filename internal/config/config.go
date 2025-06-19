package config

import (
	"os"
)

// Config holds all configuration for the application
type Config struct {
	Port      string
	InfuraURL string
}

// Load creates a new configuration instance with environment variables
func Load() *Config {
	return &Config{
		Port:      getEnv("PORT", "8080"),
		InfuraURL: getEnv("INFURA_URL", "https://mainnet.infura.io/v3/YOUR_API_KEY"),
	}
}

// getEnv retrieves environment variable with fallback to default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
