package business

import "errors"

// ErrTransactionAlreadyExists -
var ErrTransactionAlreadyExists = errors.New("order way already paid by a user")
