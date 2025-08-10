package configs

import (
	"fmt"
	"net/url"
)

type MySQLConfig struct {
	Host     string `env:"MYSQL_HOST" validate:"required"`
	Port     string `env:"MYSQL_PORT" validate:"required"`
	Name     string `env:"MYSQL_NAME" validate:"required"`
	Username string `env:"MYSQL_USER" validate:"required"`
	Password string `env:"MYSQL_PASS" validate:"required"`
}

func (c MySQLConfig) ToGormURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.Name,
	)
}

func (c MySQLConfig) ToDbmateURL() string {
	return fmt.Sprintf("mysql://%s:%s@%s:%s/%s",
		c.Username,
		url.QueryEscape(c.Password),
		c.Host,
		c.Port,
		c.Name,
	)
}
