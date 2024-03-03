package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"payments/internal/domain/entities"
	"time"
)

type MyBankPaymentRepository struct {
	tableName string
	database  *mongo.Database
}

func (r MyBankPaymentRepository) Pay(ctx context.Context, originAccount string, destinationAccount string, amount float64) (*entities.Transaction, error) {
	t := &entities.Transaction{
		Amount:             amount,
		OriginAccount:      originAccount,
		DestinationAccount: destinationAccount,
		CreatedAt:          time.Now(),
		Type:               "PAY",
	}
	collection := r.database.Collection(r.tableName)
	result, err := collection.InsertOne(ctx, t)
	if err != nil {
		return nil, err
	}
	t.Id = result.InsertedID.(primitive.ObjectID).Hex()
	return t, nil
}

func (r MyBankPaymentRepository) Refund(ctx context.Context, transactionId string) error {
	transaction, err := r.GetTransaction(ctx, transactionId)
	if err != nil {
		return err
	}

	t := &entities.Transaction{
		Amount:             transaction.Amount,
		OriginAccount:      transaction.OriginAccount,
		DestinationAccount: transaction.DestinationAccount,
		CreatedAt:          time.Now(),
		Type:               "REFUND",
	}

	collection := r.database.Collection(r.tableName)
	_, err = collection.InsertOne(ctx, t)
	if err != nil {
		return err
	}
	return nil

}

func (r MyBankPaymentRepository) History(ctx context.Context) ([]entities.Customer, error) {
	//TODO implement me
	panic("implement me")
}

func (r MyBankPaymentRepository) GetTransaction(ctx context.Context, transactionId string) (*entities.Transaction, error) {
	var result entities.Transaction
	collection := r.database.Collection(r.tableName)
	objectId, err := primitive.ObjectIDFromHex(transactionId)
	if err != nil {
		return nil, err
	}
	filter := bson.M{
		"_id": objectId,
	}
	err = collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func NewMyBankPaymentRepository(db *mongo.Database) *MyBankPaymentRepository {
	return &MyBankPaymentRepository{
		database:  db,
		tableName: "transactions",
	}
}
