package app

import (
	payment_api "github.com/restinbass/payment-service/internal/api/payment/v1"
	payment_v1 "github.com/restinbass/payment-service/pkg/proto/payment/v1"
)

// APIs -
type APIs struct {
	Payment payment_v1.PaymentServiceServer
}

// InitAPIs -
func InitAPIs(services Services) APIs {
	return APIs{
		Payment: payment_api.New(services.Payment),
	}
}
