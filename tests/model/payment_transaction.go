package test_model

import "github.com/google/uuid"

type (
	// PaymentTransaction -
	PaymentTransaction struct {
		TransactionID uuid.UUID `db:"id"`
		OrderID       uuid.UUID `db:"order_id"`
		UserID        uuid.UUID `db:"user_id"`
		PaymentMethod int32     `db:"payment_method"`
	}
)
