package domain

import "github.com/google/uuid"

type (
	// PaymentTransaction -
	PaymentTransaction struct {
		TransactionID uuid.UUID
		OrderID       uuid.UUID
		UserID        uuid.UUID
		PaymentMethod PaymentMethod
	}
)
