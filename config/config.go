package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	ApiPort     string `envconfig:"API_PORT" default:"8088"`
	ApiLogPath  string `envconfig:"API_LOG_PATH"`
	ApiLogLevel string `envconfig:"API_LOG_LEVEL" default:"debug"`
}

func ConfInitialize() (*Config, error) {
	c := new(Config)

	if err := envconfig.Process("api", c); err != nil {
		log.Println("[Config] config failed read config, err: ", err)
		return nil, err
	}

	return c, nil
}
