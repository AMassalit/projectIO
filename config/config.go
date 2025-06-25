package config

import (
	"fmt"

	"github.com/caarlos0/env/v9"
)

type Config struct {
	Log
	HTTP
}

type Log struct {
	Level string `env:"LOG_LEVEL" envDefault:"debug"`
}

type HTTP struct {
	Port int    `env:"PORT" envDefault:"8080"`
	Host string `env:"HOST" envDefault:"localhost"`
}

func NewConfig() (*Config, error) {

	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil

}

func (h HTTP) Address() string {
	return fmt.Sprintf("%s:%d", h.Host, h.Port)
}
