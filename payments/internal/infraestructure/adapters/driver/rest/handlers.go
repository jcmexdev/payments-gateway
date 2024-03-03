package rest

import (
	swagger "github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"log"
	"payments/internal/application/ports/services"
	driven_fiber "payments/internal/infraestructure/adapters/driven/fiber"
)

type Handler struct {
	Fiber                *driven_fiber.Server
	myBankPaymentService services.IPaymentsService
}

func NewRestHandler(s services.IPaymentsService) *Handler {
	Fiber := driven_fiber.NewFiberServer()
	return &Handler{
		Fiber:                Fiber,
		myBankPaymentService: s,
	}
}

func (r *Handler) SetUpRoutes() {
	v1 := r.Fiber.Server.Group("/api/v1")
	v1.Get("/health", r.HealthCheck)
	v1.Post("/transaction", r.Transfer)
	v1.Get("/transaction/:transactionId", r.PaymentDetails)
	v1.Post("/transaction/:transactionId/refund", r.Refund)
	v1.Get("/docs/*", swagger.New(swagger.ConfigDefault))
}

func (r *Handler) Start(address string) {
	err := r.Fiber.Server.Listen(address)

	if err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}

type ServerResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Body    any    `json:"body,omitempty"`
}

func JsonResponse(ctx *fiber.Ctx, status int, message string, body any) error {
	return ctx.JSON(ServerResponse{
		Status:  status,
		Message: message,
		Body:    body,
	})
}
