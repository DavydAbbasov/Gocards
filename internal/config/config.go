package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Environment string `env:"ENVIRONMENT"`
	HTTPServer  HTTPServerConfig
}

type HTTPServerConfig struct {
	Address string        `env:"HTTP_SERVER_ADDRESS"`
	Port    string        `env:"HTTP_SERVER_PORT"`
	Timeout time.Duration `env:"HTTP_SERVER_TIMEOUT"`
}

func Load() (*Config, error) {
	var cfg Config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, fmt.Errorf("failed to read config from env: %w", err)
	}

	if cfg.HTTPServer.Timeout == 0 {
		cfg.HTTPServer.Timeout = 15 * time.Second
	}

	//to do

	return &cfg, nil
}
