package configs

type MySQLConfig struct {
	Host     string `env:"MYSQL_HOST" validate:"required"`
	Port     string `env:"MYSQL_PORT" validate:"required"`
	Name     string `env:"MYSQL_NAME" validate:"required"`
	Username string `env:"MYSQL_USER" validate:"required"`
	Password string `env:"MYSQL_PASS" validate:"required"`
}
