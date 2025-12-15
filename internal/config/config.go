package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Config holds the application's configuration.
type Config struct {
	DBSource string `envconfig:"DB_SOURCE" required:"true"`
}

// Load loads the configuration from environment variables.
// It also loads from a .env file if it exists.
func Load() (*Config, error) {
	// Load .env file, but ignore errors if it doesn't exist.
	_ = godotenv.Load()

	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
