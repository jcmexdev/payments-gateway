package repositories

import (
	"context"
	"payments/internal/domain/entities"
)

type IPaymentsRepository interface {
	Pay(ctx context.Context, originAccount string, destinationAccount string, amount float64) (*entities.Transaction, error)
	Refund(ctx context.Context, transactionId string) error
	History(ctx context.Context) ([]entities.Customer, error)
	GetTransaction(ctx context.Context, transactionId string) (*entities.Transaction, error)
}
