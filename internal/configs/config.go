package configs

import (
	"encoding/json"
	"errors"
	"go-boilerplate/internal/constants"
	"go-boilerplate/pkg/customerror"
	"go-boilerplate/pkg/customvalidator"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	Port               int                `env:"PORT" validate:"required"`
	Environment        constants.EnvValue `env:"ENVIRONMENT" validate:"oneof=production development"`
	CORSAllowedOrigins []string           `env:"CORS_ALLOWED_ORIGINS" envSeparator:","`
	LogsDir            string             `env:"LOGS_DIR" validate:"required"`

	MySQL MySQLConfig `validate:"required"`
}

var defaultConfig Config

func Load() (err error) {
	if err = godotenv.Load(); err != nil {
		return
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
