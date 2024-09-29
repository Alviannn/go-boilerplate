package configs

import (
	"go-boilerplate/internal/constants"
	"go-boilerplate/pkg/customvalidator"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	Port               int                `env:"PORT"`
	Environment        constants.EnvValue `env:"ENVIRONMENT" validate:"oneof=production development"`
	CORSAllowedOrigins []string           `env:"CORS_ALLOWED_ORIGINS" envSeparator:","`

	MySQL MySQLConfig
}

var defaultConfig Config

func Load() (err error) {
	if err = godotenv.Load(); err != nil {
		return
	}
	if err = env.Parse(&defaultConfig); err != nil {
		return
	}

	validator := customvalidator.New()
	err = validator.Validate(&defaultConfig)
	return
}

func Default() Config {
	return defaultConfig
}
