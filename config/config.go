package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	ApiPort string `envconfig:"API_PORT" default:"8088"`
}

func ConfInitialize() (*Config, error) {
	c := new(Config)

	if err := envconfig.Process("api", c); err != nil {
		log.Println("[Config] config failed read config, err: ", err)
		return nil, err
	}

	return c, nil
}
