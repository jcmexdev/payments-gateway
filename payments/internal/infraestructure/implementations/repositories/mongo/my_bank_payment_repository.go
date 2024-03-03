package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"payments/internal/domain/entities"
	"time"
)

type MyBankPaymentRepository struct {
	tableName string
	database  *mongo.Database
}

func (r MyBankPaymentRepository) Pay(ctx context.Context, originAccount string, destinationAccount string, amount float64) error {

	example := &entities.Transaction{
		Amount:             amount,
		OriginAccount:      originAccount,
		DestinationAccount: destinationAccount,
		CreatedAt:          time.Now(),
		Type:               "PAY",
	}
	collection := r.database.Collection(r.tableName)
	_, err := collection.InsertOne(ctx, example)
	if err != nil {
		return err
	}
	return nil
}

func (r MyBankPaymentRepository) Refund(ctx context.Context, transactionId string) error {
	//TODO implement me
	panic("implement me")
}

func (r MyBankPaymentRepository) History(ctx context.Context) ([]entities.Customer, error) {
	//TODO implement me
	panic("implement me")
}

func NewMyBankPaymentRepository(db *mongo.Database) *MyBankPaymentRepository {
	return &MyBankPaymentRepository{
		database:  db,
		tableName: "transactions",
	}
}
