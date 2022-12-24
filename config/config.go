package config

import (
	"os"

	"github.com/Netflix/go-env"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

// Config represents application configuration which can also be set
// via environment variable during runtime.
type Config struct {
	App struct {
		APP struct {
			URL       string `default:"http://localhost" env:"APP_URL"`
			Port      uint   `default:"8080" env:"APP_PORT"`
			Timezone  string `default:"UTC" env:"APP_TIMEZONE"`
			DebugMode bool   `default:"false" env:"APP_DEBUG_MODE"`
		}
		DevMode bool `default:"false" env:"DEV"`
	}
}

// LoadConfig loads the configuration from `.env` file in the same
// directory as the application and populate the Config accordingly.
func LoadConfig() (*Config, error) {
	configLogger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	envFile := os.Getenv("ENV_FILE")
	if envFile == "" {
		envFile = ".env"
	}

	if err := godotenv.Load(envFile); err != nil {
		configLogger.Info("Error loading dot env file", zap.String("path", envFile))
	}

	var cfg Config
	_, err = env.UnmarshalFromEnviron(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, err
}

// LoadTestConfig ...
func LoadTestConfig() (*Config, error) {
	config, err := LoadConfig()
	if err != nil {
		return nil, err
	}

	return config, err
}
