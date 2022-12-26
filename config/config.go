package config

import (
	"log"
	"os"

	"github.com/jinzhu/configor"
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
	// Since the application logger depends on the config, we will use a dedicated logger here.
	configLogger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	err = loadEnv()
	if err == nil {
		configLogger.Info("Env loaded")
	}

	config := Config{}
	err = configor.Load(&config)

	if err != nil {
		configLogger.Fatal("Error loading config", zap.Error(err))
		os.Exit(1)
	}

	return &config, err
}

// LoadTestConfig ...
func LoadTestConfig() (*Config, error) {
	config, err := LoadConfig()
	if err != nil {
		return nil, err
	}

	return config, err
}
