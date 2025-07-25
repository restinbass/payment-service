package payment_service

import (
	"github.com/restinbass/payment-service/internal/repository"
	"github.com/restinbass/payment-service/internal/service"
)

var _ service.PaymentTransactionService = (*serviceImpl)(nil)

type serviceImpl struct {
	paymentRepo repository.PaymentTransactionRepository
}

// New -
func New(paymentRepo repository.PaymentTransactionRepository) *serviceImpl {
	return &serviceImpl{
		paymentRepo: paymentRepo,
	}
}
