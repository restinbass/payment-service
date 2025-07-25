package business

type PaymentMethod int32

// PaymentMethod -
const (
	PaymentMethodUncpecified PaymentMethod = iota
	PaymentMethodCard
	PaymentMethodSBP
	PaymentMethodCreditCard
	PaymentMethodInvestorMoney
)
