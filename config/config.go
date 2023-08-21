package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	HttpInfo struct {
		Port string `envconfig:"API_PORT" default:"8088"`
	}

	LogInfo struct {
		LogPath  string `envconfig:"API_LOG_PATH"`
		LogLevel string `envconfig:"API_LOG_LEVEL" default:"debug"`
	}

	MySqlInfo struct {
		Host   string `envconfig:"API_DB_HOST" default:"localhost"`
		Port   string `envconfig:"API_DB_PORT" default:"3306"`
		DbName string `envconfig:"API_DB_NAME" default:"catchreview"`
		User   string `envconfig:"API_DB_USER" default:"root"`
		Pass   string `envconfig:"API_DB_PASSWORD" default:"1234"`
	}
}

func ConfInitialize() (*Config, error) {
	c := new(Config)

	if err := envconfig.Process("api", c); err != nil {
		log.Println("[Config] config failed read config, err: ", err)
		return nil, err
	}

	return c, nil
}
