package configs

type MySQLConfig struct {
	Host     string `env:"MYSQL_HOST"`
	Port     string `env:"MYSQL_PORT"`
	Name     string `env:"MYSQL_NAME"`
	Username string `env:"MYSQL_USER"`
	Password string `env:"MYSQL_PASS"`
}
