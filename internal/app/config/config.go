package config

import (
	"time"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	Server  ServerConfig
	AppName string
	Version string
}

func New(appName, version string) (Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return Config{}, err
	}
	cfg.AppName = appName
	cfg.Version = version
	return cfg, nil
}

type ServerConfig struct {
	Port           string        `env:"HTTP_PORT"             envDefault:"8080"`
	WriteTimeout   time.Duration `env:"HTTP_WRITE_TIMEOUT"    envDefault:"20s"`
	ReadTimeout    time.Duration `env:"HTTP_READ_TIMEOUT"     envDefault:"15s"`
	IdleTimeout    time.Duration `env:"HTTP_IDLE_TIMEOUT"     envDefault:"30s"`
	MaxHeaderBytes int           `env:"HTTP_MAX_HEADER_BYTES" envDefault:"1048576"`
}
