package env_config

import (
	"github.com/caarlos0/env/v11"
)

type (
	grpcServerConfig struct {
		cfg grpcServerEnvConfig
	}

	grpcServerEnvConfig struct {
		Port int64 `env:"GRPC_PORT,required"`
	}
)

// NewGrpcServerConfig -
func NewGrpcServerConfig() (*grpcServerConfig, error) {
	cfg := grpcServerEnvConfig{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &grpcServerConfig{
		cfg: cfg,
	}, nil
}

// Port -
func (c *grpcServerConfig) Port() int64 {
	return c.cfg.Port
}
