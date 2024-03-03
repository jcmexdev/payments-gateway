package main

import (
	"context"
	"fmt"
	"payments/config"
	mongodb "payments/internal/infraestructure/adapters/driven/mongo"
	"payments/internal/infraestructure/adapters/driver/rest"
	"payments/internal/infraestructure/implementations/integrations/mybank"
	mongo_repository "payments/internal/infraestructure/implementations/repositories/mongo"
	"payments/internal/infraestructure/implementations/services"
)

/*
curl "http://localhost:3000/deposit?accountNumber=1111-2222-3333-4444&amount=100"
curl "http://localhost:3000/transfer?originAccountNumber=1111-2222-3333-4444&destinationAccountNumber=1234-5678-1234-5678&amount=100"
*/

func main() {
	config.LoadEnvironment(context.Background())
	mongodb.ConnectMongoDB(context.Background())
	defer mongodb.DisconnectMongoDB(context.Background())

	myBankRepository := mongo_repository.NewMyBankPaymentRepository(mongodb.GetDatabase())
	myBankIntegration := mybank.NewMyBankIntegration()
	services.NewMyBankPaymentService(myBankRepository, myBankIntegration)
	restServer := rest.NewRestHandler(services.NewMyBankPaymentService(myBankRepository, myBankIntegration))
	restServer.SetUpRoutes()
	address := fmt.Sprintf("%s:%s", config.Conf.AppBaseUrl, config.Conf.AppPort)
	restServer.Start(address)
}
