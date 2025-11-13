package config

import (
	"crypto/tls"
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Environment string `env:"ENVIRONMENT"`
	Redis       RedisConfig
	Database    DatabaseConfig
	HTTPServer  HTTPServerConfig
}

type HTTPServerConfig struct {
	Address string        `env:"HTTP_SERVER_ADDRESS"`
	Port    string        `env:"HTTP_SERVER_PORT"`
	Timeout time.Duration `env:"HTTP_SERVER_TIMEOUT"`
}
type RedisConfig struct {
	Addr        string        `env:"REDIS_ADDR"`
	Password    string        `env:"REDIS_PASSWORD"`
	DB          int           `env:"REDIS_DB"`
	DialTimeout time.Duration `env:"REDIS_DIAL_TIMEOUT"`
	UseTLS      bool          `env:"REDIS_USER_TLS"`
	Prefix      string        `env:"REDIS_PREFIX"`
	TLSConfig   *tls.Config   `env:"-"`
}
type DatabaseConfig struct {
	UserName string `env:"DATABASE_USER_NAME"`
	Addr     string `env:"DATABASE_ADDRESS"`
	Port     string `env:"DATABASE_PORT"`
	DBName   string `env:"DATABASE_NAME"`
	Password string `env:"DATABASE_PASSWORD"`
	SSLMODE  string `env:"DATABASE_SSLMODE"`
}

func Load() (*Config, error) {
	var cfg Config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, fmt.Errorf("failed to read config from env: %w", err)
	}

	//time
	if cfg.HTTPServer.Timeout == 0 {
		cfg.HTTPServer.Timeout = 15 * time.Second
	}
	if cfg.Redis.DialTimeout == 0 {
		cfg.Redis.DialTimeout = 5 * time.Second
	}

	return &cfg, nil
}
