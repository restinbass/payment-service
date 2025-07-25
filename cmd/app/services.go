package app

import (
	"github.com/restinbass/payment-service/internal/service"
	payment_service "github.com/restinbass/payment-service/internal/service/payment"
)

// Services -
type Services struct {
	Payment service.PaymentTransactionService
}

// InitServices -
func InitServices(repositories Repositories) Services {
	return Services{
		Payment: payment_service.New(repositories.Payment),
	}
}
