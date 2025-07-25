package payment_repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/restinbass/payment-service/internal/repository"
)

var _ repository.PaymentTransactionRepository = (*repoImpl)(nil)

type repoImpl struct {
	db *pgxpool.Pool
}

// New -
func New(db *pgxpool.Pool) *repoImpl {
	return &repoImpl{
		db: db,
	}
}
