package payment_api

import (
	"github.com/restinbass/payment-service/internal/service"
	payment_v1 "github.com/restinbass/payment-service/pkg/proto/payment/v1"
)

type apiImpl struct {
	payment_v1.UnimplementedPaymentServiceServer

	paymentService service.PaymentTransactionService
}

// New -
func New(paymentService service.PaymentTransactionService) *apiImpl {
	return &apiImpl{
		paymentService: paymentService,
	}
}
