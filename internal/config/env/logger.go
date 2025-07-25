package env_config

import (
	"github.com/caarlos0/env/v11"
	"github.com/restinbass/platform-libs/pkg/logger"
)

type (
	loggerConfig struct {
		cfg loggerEnvConfig
	}

	loggerEnvConfig struct {
		LogLevel string `env:"LOG_LEVEL,required"`
		AsJSON   bool   `env:"LOG_AS_JSON,required"`
	}
)

// NewLoggerConfig -
func NewLoggerConfig() (*loggerConfig, error) {
	cfg := loggerEnvConfig{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &loggerConfig{
		cfg: cfg,
	}, nil
}

// LogLevel -
func (c *loggerConfig) LogLevel() logger.LogLevel {
	return logger.LogLevel(c.cfg.LogLevel)
}

// AsJSON -
func (c *loggerConfig) AsJSON() bool {
	return c.cfg.AsJSON
}
