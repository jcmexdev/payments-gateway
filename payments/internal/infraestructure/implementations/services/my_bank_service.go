package services

import (
	"context"
	"payments/internal/application/ports/integrations"
	"payments/internal/application/ports/repositories"
	"payments/internal/domain/entities"
)

type MyBankPaymentsService struct {
	ctx         context.Context
	repo        repositories.IPaymentsRepository
	integration integrations.IBankIntegration
}

func (p MyBankPaymentsService) Pay(ctx context.Context, originAccount string, destinationAccount string, amount float64) (*entities.Transaction, error) {
	_, err := p.integration.Transfer(originAccount, destinationAccount, amount)
	if err != nil {
		return nil, err
	}
	transaction, err := p.repo.Pay(ctx, originAccount, destinationAccount, amount)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (p MyBankPaymentsService) Refund(ctx context.Context, transactionId string) error {
	transaction, err := p.repo.GetTransaction(ctx, transactionId)
	if err != nil {
		return err
	}
	err = p.integration.Refund(*transaction)
	if err != nil {
		return err
	}
	err = p.repo.Refund(ctx, transactionId)
	if err != nil {
		return err
	}
	return nil
}

func (p MyBankPaymentsService) History(ctx context.Context) ([]entities.Customer, error) {
	history, err := p.repo.History(nil)
	if err != nil {
		return nil, err
	}
	return history, nil
}

func (p MyBankPaymentsService) GetTransaction(ctx context.Context, transactionId string) (*entities.Transaction, error) {
	transaction, err := p.repo.GetTransaction(ctx, transactionId)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func NewMyBankPaymentService(repo repositories.IPaymentsRepository, myBankIntegration integrations.IBankIntegration) *MyBankPaymentsService {
	return &MyBankPaymentsService{repo: repo, integration: myBankIntegration}
}
