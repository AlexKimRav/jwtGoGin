package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBUsername string `mapstructure:"MYSQL_USER"`
	DBPassword string `mapstructure:"MYSQL_PASSWORD"`
	DBName     string `mapstructure:"MYSQL_DB"`

	TokenSecret    string        `mapstructure:"my-secret"`
	TokenExpiresIn time.Duration `mapstructure:"TOKEN_EXPIRED_IN"`
	TokenMaxAge    int           `mapstructure:"TOKEN_MAXAGE"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return config, err
}
