package rest

import (
	"github.com/gofiber/fiber/v2"
	"payments/internal/infraestructure/adapters/dtos/request"
)

func (r *Handler) HealthCheck(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"status": "ok"})
}

func (r *Handler) Transfer(ctx *fiber.Ctx) error {
	var requestBody request.TransferRequest

	if err := ctx.BodyParser(&requestBody); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	}

	originAccount := requestBody.OriginAccount
	destinationAccount := requestBody.DestinationAccount
	amount := requestBody.Amount

	transaction, err := r.myBankPaymentService.Pay(ctx.Context(), originAccount, destinationAccount, amount)

	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, err.Error())
	}

	return JsonResponse(ctx, fiber.StatusCreated, "Transaction successful", transaction)
}

func (r *Handler) Refund(ctx *fiber.Ctx) error {
	transactionId := ctx.Params("transactionId")
	if transactionId == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid transaction id")
	}

	err := r.myBankPaymentService.Refund(ctx.Context(), transactionId)

	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, err.Error())
	}

	return JsonResponse(ctx, fiber.StatusOK, "Refund successful", nil)
}

func (r *Handler) PaymentDetails(ctx *fiber.Ctx) error {
	transactionId := ctx.Params("transactionId")

	transaction, err := r.myBankPaymentService.GetTransaction(ctx.Context(), transactionId)

	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, err.Error())
	}

	return JsonResponse(ctx, fiber.StatusOK, "Transaction details", transaction)
}
