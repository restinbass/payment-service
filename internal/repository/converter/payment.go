package domain_converter

import (
	"github.com/google/uuid"
	domain "github.com/restinbass/payment-service/internal/repository/model"
	business "github.com/restinbass/payment-service/internal/service/model"
)

// CreatePaymentTransactionRequestToTransaction -
func CreatePaymentTransactionRequestToTransaction(req business.CreatePaymentTransactionRequest) domain.PaymentTransaction {
	return domain.PaymentTransaction{
		TransactionID: uuid.New(),
		OrderID:       req.OrderID,
		UserID:        req.UserID,
		PaymentMethod: domain.PaymentMethod(req.PaymentMethod),
	}
}

// PaymentTransactionToBusiness -
func PaymentTransactionToBusiness(t domain.PaymentTransaction) business.CreatePaymentTransactionResponse {
	return business.CreatePaymentTransactionResponse{
		PaymentTransaction: business.PaymentTransaction{
			OrderID:       t.OrderID,
			UserID:        t.UserID,
			PaymentMethod: business.PaymentMethod(t.PaymentMethod),
			TransactionID: t.TransactionID,
		},
	}
}
