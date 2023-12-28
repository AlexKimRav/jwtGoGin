package config

import "time"

type Config struct {
	DBHost     string `mapstructure:"MYSQL_HOST"`
	DBUsername string `mapstructure:"MYSQL_USER"`
	DBPassword string `mapstructure:"MYSQL_PASSWORD"`
	DBName     string `mapstructure:"MYSQL_DB"`
	DBPort     string `mapstructure:"MYSQL_HOST"`
	Serverpost string `mapstructure:"PORT"`

	TokenSecret    string        `mapstructure:"my-secret"`
	TokenExpiresIn time.Duration `mapstructure:"TOKEN_EXPIRED_IN"`
	TokenMaxAge    int           `mapstructure:"TOKEN_MAXAGE"`
}
