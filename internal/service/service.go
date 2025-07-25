package service

import (
	"context"

	business "github.com/restinbass/payment-service/internal/service/model"
)

// PaymentTransactionService -
type PaymentTransactionService interface {
	Create(ctx context.Context, req business.CreatePaymentTransactionRequest) (business.CreatePaymentTransactionResponse, error)
}
