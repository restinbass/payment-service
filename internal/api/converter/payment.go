package api_converter

import (
	"github.com/google/uuid"
	business "github.com/restinbass/payment-service/internal/service/model"
	payment_v1 "github.com/restinbass/payment-service/pkg/proto/payment/v1"
)

func PayOrderRequestToBusiness(req *payment_v1.PayOrderRequest) business.CreatePaymentTransactionRequest {
	return business.CreatePaymentTransactionRequest{
		OrderID:       uuid.MustParse(req.GetOrderUuid()),
		UserID:        uuid.MustParse(req.GetUserUuid()),
		PaymentMethod: business.PaymentMethod(req.GetPaymentMethod()),
	}
}

func CratePaymentTransactionResponseToAPI(resp business.CreatePaymentTransactionResponse) *payment_v1.PayOrderResponse {
	return &payment_v1.PayOrderResponse{
		TransactionUuid: resp.PaymentTransaction.TransactionID.String(),
	}
}
