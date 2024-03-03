package integrations

import (
	"payments/internal/domain/entities"
	"payments/internal/infraestructure/adapters/dtos"
)

type IBankIntegration interface {
	Transfer(originAccount string, destinationAccount string, amount float64) (response *dtos.MyBankResponse, err error)
	Refund(transaction entities.Transaction) error
}
