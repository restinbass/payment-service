package payment_service_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/restinbass/payment-service/internal/repository"
	"github.com/restinbass/payment-service/internal/repository/mocks"
	business "github.com/restinbass/payment-service/internal/service/model"
	payment_service "github.com/restinbass/payment-service/internal/service/payment"
	"github.com/restinbass/platform-libs/pkg/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreatePaymentTransaction(t *testing.T) {
	type testCase struct {
		name         string
		req          business.CreatePaymentTransactionRequest
		repoMock     func(t *testing.T, tt testCase) repository.PaymentTransactionRepository
		expectedResp business.CreatePaymentTransactionResponse
		expectedErr  error
	}

	tests := []testCase{
		{
			name: "valid payment create request",
			req: business.CreatePaymentTransactionRequest{
				OrderID:       uuid.MustParse("13cd4203-dff8-49e9-bc95-47c36d35438e"),
				UserID:        uuid.MustParse("a19e3612-4f8e-4984-8ce1-3177e0a5ec01"),
				PaymentMethod: business.PaymentMethodCard,
			},
			repoMock: func(t *testing.T, tt testCase) repository.PaymentTransactionRepository {
				m := mocks.NewPaymentTransactionRepository(t)
				m.EXPECT().Create(mock.Anything, tt.req).
					Return(business.CreatePaymentTransactionResponse{
						PaymentTransaction: business.PaymentTransaction{
							TransactionID: uuid.MustParse("3edee576-fa7c-4a6a-abd2-cdfb528d057e"),
							OrderID:       uuid.MustParse("13cd4203-dff8-49e9-bc95-47c36d35438e"),
							UserID:        uuid.MustParse("a19e3612-4f8e-4984-8ce1-3177e0a5ec01"),
							PaymentMethod: business.PaymentMethodCard,
						},
					}, nil).
					Once()

				return m
			},
			expectedResp: business.CreatePaymentTransactionResponse{
				PaymentTransaction: business.PaymentTransaction{
					TransactionID: uuid.MustParse("3edee576-fa7c-4a6a-abd2-cdfb528d057e"),
					OrderID:       uuid.MustParse("13cd4203-dff8-49e9-bc95-47c36d35438e"),
					UserID:        uuid.MustParse("a19e3612-4f8e-4984-8ce1-3177e0a5ec01"),
					PaymentMethod: business.PaymentMethodCard,
				},
			},
			expectedErr: nil,
		},
		{
			name: "invalid payment create request (user already paid)",
			req: business.CreatePaymentTransactionRequest{
				OrderID:       uuid.MustParse("13cd4203-dff8-49e9-bc95-47c36d35438e"),
				UserID:        uuid.MustParse("a19e3612-4f8e-4984-8ce1-3177e0a5ec01"),
				PaymentMethod: business.PaymentMethodCard,
			},
			repoMock: func(t *testing.T, tt testCase) repository.PaymentTransactionRepository {
				m := mocks.NewPaymentTransactionRepository(t)
				m.EXPECT().Create(mock.Anything, tt.req).
					Return(business.CreatePaymentTransactionResponse{}, business.ErrTransactionAlreadyExists).
					Once()

				return m
			},
			expectedResp: business.CreatePaymentTransactionResponse{},
			expectedErr:  business.ErrTransactionAlreadyExists,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger.SetNopLogger()

			paymentService := payment_service.New(tt.repoMock(t, tt))
			resp, err := paymentService.Create(context.Background(), tt.req)

			assert.ErrorIs(t, err, tt.expectedErr)
			assert.Equal(t, tt.expectedResp, resp)
		})
	}
}
