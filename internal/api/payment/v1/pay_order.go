package payment_api

import (
	"context"
	"errors"

	api_converter "github.com/restinbass/payment-service/internal/api/converter"
	business "github.com/restinbass/payment-service/internal/service/model"
	payment_v1 "github.com/restinbass/payment-service/pkg/proto/payment/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *apiImpl) PayOrder(ctx context.Context, req *payment_v1.PayOrderRequest) (*payment_v1.PayOrderResponse, error) {
	resp, err := a.paymentService.Create(ctx, api_converter.PayOrderRequestToBusiness(req))
	if err != nil {
		if errors.Is(err, business.ErrTransactionAlreadyExists) {
			return nil, status.Errorf(codes.AlreadyExists, "user already paid for this order")
		}
		return nil, err
	}

	return api_converter.CratePaymentTransactionResponseToAPI(resp), nil
}
