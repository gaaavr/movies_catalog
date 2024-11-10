package config

import (
	"fmt"

	"github.com/caarlos0/env/v9"
)

type Config struct {
	Postgres PostgresConfig
	Server   ServerConfig
	Email    EmailConfig
}

type PostgresConfig struct {
	Host     string `env:"POSTGRES_HOST" envDefault:"127.0.0.1"`
	Port     string `env:"POSTGRES_PORT" envDefault:"5432"`
	Username string `env:"POSTGRES_USERNAME" envDefault:""`
	Password string `env:"POSTGRES_PASSWORD" envDefault:""`
	Database string `env:"POSTGRES_DATABASE" envDefault:""`
	SslMode  string `env:"POSTGRES_SSL_MODE" envDefault:"disable"`
}

type ServerConfig struct {
	Host string `env:"SERVICE_HOST" envDefault:"0.0.0.0"`
	Port string `env:"SERVICE_PORT" envDefault:"8080"`
}

type EmailConfig struct {
	Username string `env:"EMAIL_USERNAME"`
	AppPass  string `env:"EMAIL_SENDER_APP_PASS"`
	Server   string `env:"EMAIL_SERVER_NAME"`
	Port     int    `env:"EMAIL_PORT"`
}

func New() (*Config, error) {
	var cfg Config

	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	return &cfg, nil
}
