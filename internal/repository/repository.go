package repository

import (
	"context"

	business "github.com/restinbass/payment-service/internal/service/model"
)

// PaymentTransactionRepository -
type PaymentTransactionRepository interface {
	Create(ctx context.Context, req business.CreatePaymentTransactionRequest) (business.CreatePaymentTransactionResponse, error)
}
