package configs

import (
	"encoding/json"
	"errors"
	"go-boilerplate/internal/constants"
	"go-boilerplate/pkg/customerror"
	"go-boilerplate/pkg/customvalidator"
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	Port               int                `env:"PORT"`
	Environment        constants.EnvValue `env:"ENVIRONMENT" validate:"oneof=production development"`
	CORSAllowedOrigins []string           `env:"CORS_ALLOWED_ORIGINS" envSeparator:","`
	LogsDir            string             `env:"LOGS_DIR" validate:"required"`

	MySQL MySQLConfig `validate:"required"`
}

func (c *Config) Validate() (err error) {
	if c.Port == 0 {
		c.Port = 5000
	}
	return
}

var defaultConfig Config

func Load() (err error) {
	envFilePath := ".env"

	if _, err = os.Stat(envFilePath); err == nil {
		if err = godotenv.Load(envFilePath); err != nil {
			return
		}
	}
	if err = env.Parse(&defaultConfig); err != nil {
		return
	}

	if err = customvalidator.New().Validate(&defaultConfig); err != nil {
		customErr := customerror.New().
			WithSourceError(err).
			WithMessage("Failed to validate config")

		errJson, _ := json.Marshal(customErr.ToJSON())
		err = errors.New(string(errJson))
		return
	}

	return
}

func Default() Config {
	return defaultConfig
}
