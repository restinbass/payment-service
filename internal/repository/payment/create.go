package payment_repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	domain_converter "github.com/restinbass/payment-service/internal/repository/converter"
	domain "github.com/restinbass/payment-service/internal/repository/model"
	business "github.com/restinbass/payment-service/internal/service/model"
	"github.com/restinbass/platform-libs/pkg/logger"
	"go.uber.org/zap"
)

// Create -
func (r *repoImpl) Create(ctx context.Context, req business.CreatePaymentTransactionRequest) (business.CreatePaymentTransactionResponse, error) {
	transaction := domain_converter.CreatePaymentTransactionRequestToTransaction(req)

	exists, err := r.checkTransactionAlreadyExists(ctx, transaction)
	if err != nil {
		return business.CreatePaymentTransactionResponse{}, err
	}

	if exists {
		return business.CreatePaymentTransactionResponse{}, business.ErrTransactionAlreadyExists
	}

	qb := sq.Insert("payment_transactions").
		PlaceholderFormat(sq.Dollar).
		Columns("id", "order_id", "user_id", "payment_method").
		Values(
			transaction.TransactionID.String(),
			transaction.OrderID.String(),
			transaction.UserID.String(),
			transaction.PaymentMethod,
		).
		Suffix("RETURNING id")

	query, args, err := qb.ToSql()
	if err != nil {
		logger.Error(ctx, "failed to ToSql", zap.Error(err))
		return business.CreatePaymentTransactionResponse{}, err
	}

	if _, err := r.db.Exec(ctx, query, args...); err != nil {
		logger.Error(ctx, "failed to db.Exec", zap.Error(err))
		return business.CreatePaymentTransactionResponse{}, err
	}

	return domain_converter.PaymentTransactionToBusiness(transaction), nil
}

func (r *repoImpl) checkTransactionAlreadyExists(ctx context.Context, transaction domain.PaymentTransaction) (bool, error) {
	qb := sq.Select("1").
		PlaceholderFormat(sq.Dollar).
		Prefix("SELECT EXISTS (").
		From("payment_transactions").
		Where(sq.Eq{"order_id": transaction.OrderID}).
		Where(sq.Eq{"user_id": transaction.UserID}).
		Suffix(")")

	query, args, err := qb.ToSql()
	if err != nil {
		logger.Error(ctx, "failed to ToSql", zap.Error(err))
		return false, err
	}

	var exists bool
	if err := r.db.QueryRow(ctx, query, args...).Scan(&exists); err != nil {
		logger.Error(ctx, "failed to QueryRow", zap.Error(err))
		return false, err
	}

	return exists, nil
}
