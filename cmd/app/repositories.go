package app

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/restinbass/payment-service/internal/config"
	"github.com/restinbass/payment-service/internal/repository"
	payment_repository "github.com/restinbass/payment-service/internal/repository/payment"
	"github.com/restinbass/platform-libs/pkg/closer"
	"github.com/restinbass/platform-libs/pkg/logger"
	"go.uber.org/zap"
)

// Repositories -
type Repositories struct {
	Payment repository.PaymentTransactionRepository
}

// InitRepositories -
func InitRepositories(ctx context.Context, cfg config.PostgresConfig) Repositories {
	pool, err := pgxpool.New(ctx, cfg.URI())
	if err != nil {
		logger.Fatal(ctx, "can not pgxpool.New", zap.Error(err))
	}
	closer.AddNamed("postgres pgxpool", func(ctx context.Context) error {
		pool.Close()
		return nil
	})

	if err := pool.Ping(ctx); err != nil {
		logger.Fatal(ctx, "failed to ping pgxpool", zap.Error(err))
	}

	return Repositories{
		Payment: payment_repository.New(pool),
	}
}
