package payment_service

import (
	"context"

	business "github.com/restinbass/payment-service/internal/service/model"
	"github.com/restinbass/platform-libs/pkg/logger"
	"go.uber.org/zap"
)

// Create -
func (s *serviceImpl) Create(ctx context.Context, req business.CreatePaymentTransactionRequest) (business.CreatePaymentTransactionResponse, error) {
	resp, err := s.paymentRepo.Create(ctx, req)
	if err != nil {
		logger.Error(
			ctx,
			"user already paid for this order",
			zap.String("order_id", req.OrderID.String()),
			zap.String("user_id", req.UserID.String()),
		)
		return business.CreatePaymentTransactionResponse{}, err
	}

	return resp, nil
}
