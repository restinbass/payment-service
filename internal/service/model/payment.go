package business

import "github.com/google/uuid"

type (
	// CreatePaymentTransactionRequest -
	CreatePaymentTransactionRequest struct {
		OrderID       uuid.UUID
		UserID        uuid.UUID
		PaymentMethod PaymentMethod
	}

	// CreatePaymentTransactionResponse -
	CreatePaymentTransactionResponse struct {
		PaymentTransaction PaymentTransaction
	}

	// PaymentTransaction -
	PaymentTransaction struct {
		TransactionID uuid.UUID
		OrderID       uuid.UUID
		UserID        uuid.UUID
		PaymentMethod PaymentMethod
	}
)
