package payment_api_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	api_converter "github.com/restinbass/payment-service/internal/api/converter"
	payment_api "github.com/restinbass/payment-service/internal/api/payment/v1"
	"github.com/restinbass/payment-service/internal/service"
	"github.com/restinbass/payment-service/internal/service/mocks"
	business "github.com/restinbass/payment-service/internal/service/model"
	payment_v1 "github.com/restinbass/payment-service/pkg/proto/payment/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestPayOrder(t *testing.T) {
	type testCase struct {
		name         string
		req          *payment_v1.PayOrderRequest
		serviceMock  func(t *testing.T, tt testCase) service.PaymentTransactionService
		expectedResp *payment_v1.PayOrderResponse
		expectedErr  error
	}

	otherServiceError := gofakeit.Error()
	tests := []testCase{
		{
			name: "valid payment create request",
			req: &payment_v1.PayOrderRequest{
				OrderUuid:     "d3db182e-381a-4840-8d22-26a7137d52eb",
				UserUuid:      "601f1a97-e618-4d0e-996c-a226f6d121cb",
				PaymentMethod: payment_v1.PayOrderRequest_PAYMENT_METHOD_CARD,
			},
			serviceMock: func(t *testing.T, tt testCase) service.PaymentTransactionService {
				m := mocks.NewPaymentTransactionService(t)
				m.EXPECT().Create(mock.Anything, api_converter.PayOrderRequestToBusiness(tt.req)).
					Return(business.CreatePaymentTransactionResponse{
						PaymentTransaction: business.PaymentTransaction{
							TransactionID: uuid.MustParse("c1358d37-8d91-4fa0-8d34-b16cb4530ba0"),
							OrderID:       uuid.MustParse("d3db182e-381a-4840-8d22-26a7137d52eb"),
							UserID:        uuid.MustParse("601f1a97-e618-4d0e-996c-a226f6d121cb"),
							PaymentMethod: business.PaymentMethodCard,
						},
					}, nil).
					Once()

				return m
			},
			expectedResp: &payment_v1.PayOrderResponse{
				TransactionUuid: "c1358d37-8d91-4fa0-8d34-b16cb4530ba0",
			},
			expectedErr: nil,
		},
		{
			name: "invalid payment create request (user already paid)",
			req: &payment_v1.PayOrderRequest{
				OrderUuid:     "d3db182e-381a-4840-8d22-26a7137d52eb",
				UserUuid:      "601f1a97-e618-4d0e-996c-a226f6d121cb",
				PaymentMethod: payment_v1.PayOrderRequest_PAYMENT_METHOD_CARD,
			},
			serviceMock: func(t *testing.T, tt testCase) service.PaymentTransactionService {
				m := mocks.NewPaymentTransactionService(t)
				m.EXPECT().Create(mock.Anything, api_converter.PayOrderRequestToBusiness(tt.req)).
					Return(business.CreatePaymentTransactionResponse{}, business.ErrTransactionAlreadyExists).
					Once()

				return m
			},
			expectedResp: nil,
			expectedErr:  status.Errorf(codes.AlreadyExists, "user already paid for this order"),
		},
		{
			name: "invalid payment create request (other service error)",
			req: &payment_v1.PayOrderRequest{
				OrderUuid:     "d3db182e-381a-4840-8d22-26a7137d52eb",
				UserUuid:      "601f1a97-e618-4d0e-996c-a226f6d121cb",
				PaymentMethod: payment_v1.PayOrderRequest_PAYMENT_METHOD_CARD,
			},
			serviceMock: func(t *testing.T, tt testCase) service.PaymentTransactionService {
				m := mocks.NewPaymentTransactionService(t)
				m.EXPECT().Create(mock.Anything, api_converter.PayOrderRequestToBusiness(tt.req)).
					Return(business.CreatePaymentTransactionResponse{}, otherServiceError).
					Once()

				return m
			},
			expectedResp: nil,
			expectedErr:  otherServiceError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			paymentAPI := payment_api.New(tt.serviceMock(t, tt))
			resp, err := paymentAPI.PayOrder(context.Background(), tt.req)

			assert.ErrorIs(t, err, tt.expectedErr)
			assert.Equal(t, tt.expectedResp, resp)
		})
	}
}
