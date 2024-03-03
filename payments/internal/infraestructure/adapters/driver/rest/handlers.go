package rest

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"payments/internal/application/ports/services"
	driven_fiber "payments/internal/infraestructure/adapters/driven/fiber"
	"strconv"
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
	v1.Get("/transfer", r.Transfer)
	v1.Get("/refund", r.Refund)
	v1.Get("/transaction/:transactionId", r.PaymentDetails)
}

func (r *Handler) Start(address string) {
	err := r.Fiber.Server.Listen(address)

	if err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}

func (r *Handler) HealthCheck(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"status": "ok"})
}

func (r *Handler) Transfer(ctx *fiber.Ctx) error {
	originAccount := ctx.Query("originAccount")
	destinationAccount := ctx.Query("destinationAccount")
	amount := ctx.Query("amount")

	amountFloat, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid amount")
	}

	transaction, err := r.myBankPaymentService.Pay(ctx.Context(), originAccount, destinationAccount, amountFloat)

	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, err.Error())
	}

	return ctx.JSON(transaction)
}

func (r *Handler) Refund(ctx *fiber.Ctx) error {
	transactionId := ctx.Params("transactionId")

	err := r.myBankPaymentService.Refund(ctx.Context(), transactionId)

	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, err.Error())
	}

	return ctx.JSON(map[string]string{"status": "ok", "message": "Refund successful"})
}

func (r *Handler) PaymentDetails(ctx *fiber.Ctx) error {
	transactionId := ctx.Params("transactionId")

	transaction, err := r.myBankPaymentService.GetTransaction(ctx.Context(), transactionId)

	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, err.Error())
	}

	return ctx.JSON(transaction)
}
