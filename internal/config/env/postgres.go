package env_config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type (
	postgresConfig struct {
		cfg postgresEnvConfig
	}

	postgresEnvConfig struct {
		User         string `env:"POSTGRES_USER,required"`
		Password     string `env:"POSTGRES_PASSWORD,required"`
		DatabaseName string `env:"POSTGRES_DB,required"`
		Host         string `env:"POSTGRES_HOST,required"`
		Port         string `env:"POSTGRES_PORT,required"`
	}
)

// NewPostgresConfig -
func NewPostgresConfig() (*postgresConfig, error) {
	cfg := postgresEnvConfig{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &postgresConfig{
		cfg: cfg,
	}, nil
}

// URI -
func (c *postgresConfig) URI() string {
	return fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=disable host=%s port=%s",
		c.cfg.User,
		c.cfg.Password,
		c.cfg.DatabaseName,
		c.cfg.Host,
		c.cfg.Port,
	)
}
