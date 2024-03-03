package mybank

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"payments/internal/infraestructure/adapters/dtos"
)

type Integration struct {
	httpClient resty.Client
	baseUrl    string
}

func (m *Integration) Transfer(originAccount string, destinationAccount string, amount float64) (response *dtos.MyBankResponse, err error) {
	var res dtos.MyBankResponse
	var errorResponse dtos.MyBankResponse
	result, err := m.httpClient.R().
		SetQueryParams(map[string]string{
			"originAccountNumber":      originAccount,
			"destinationAccountNumber": destinationAccount,
			"amount":                   fmt.Sprintf("%f", amount),
		}).
		SetHeader("Accept", "application/json").
		SetResult(&res).
		SetError(&errorResponse).
		Get(m.baseUrl + "/transfer")

	if err != nil {
		return nil, err
	}

	if result.IsSuccess() {
		println(result.String())
		return nil, nil
	}
	return nil, fiber.NewError(result.StatusCode(), errorResponse.Message)

}

func NewMyBankIntegration() *Integration {
	return &Integration{
		httpClient: *resty.New(),
		baseUrl:    "http://bank:3000",
	}
}
