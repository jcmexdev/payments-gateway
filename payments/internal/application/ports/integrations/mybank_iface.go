package integrations

import "payments/internal/infraestructure/adapters/dtos"

type IBankIntegration interface {
	Transfer(originAccount string, destinationAccount string, amount float64) (response *dtos.MyBankResponse, err error)
}
