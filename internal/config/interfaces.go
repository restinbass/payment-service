package config

import (
	"context"

	"github.com/joho/godotenv"
	env_config "github.com/restinbass/payment-service/internal/config/env"
	"github.com/restinbass/platform-libs/pkg/logger"
	"go.uber.org/zap"
)

type (
	// PostgresConfig -
	PostgresConfig interface {
		URI() string
	}

	// LoggerConfig -
	LoggerConfig interface {
		LogLevel() logger.LogLevel
		AsJSON() bool
	}

	// GrpcServerConfig -
	GrpcServerConfig interface {
		Port() int64
	}

	// Config -
	Config struct {
		PostgresConfig   PostgresConfig
		LoggerConfig     LoggerConfig
		GrpcServerConfig GrpcServerConfig
	}
)

func Load(ctx context.Context, path ...string) Config {
	if err := godotenv.Load(path...); err != nil {
		logger.Fatal(ctx, "failed to godotenv.Load()", zap.Error(err))
	}

	postgresCfg, err := env_config.NewPostgresConfig()
	if err != nil {
		logger.Fatal(ctx, "failed to init postgres config: %v", zap.Error(err))
	}

	loggerCfg, err := env_config.NewLoggerConfig()
	if err != nil {
		logger.Fatal(ctx, "failed to init logger config: %v", zap.Error(err))
	}

	grpcServerCfg, err := env_config.NewGrpcServerConfig()
	if err != nil {
		logger.Fatal(ctx, "failed to init logger config: %v", zap.Error(err))
	}

	return Config{
		PostgresConfig:   postgresCfg,
		LoggerConfig:     loggerCfg,
		GrpcServerConfig: grpcServerCfg,
	}
}
