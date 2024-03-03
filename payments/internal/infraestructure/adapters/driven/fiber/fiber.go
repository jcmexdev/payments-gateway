package fiber

import (
	"errors"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"payments/internal/infraestructure/adapters/dtos"
)

type Server struct {
	Server *fiber.App
}

func NewFiberServer() *Server {
	server := fiber.New(fiber.Config{
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}
			return ctx.Status(code).JSON(dtos.ServerErrorResponse{
				Status:  code,
				Message: err.Error(),
			})
		},
	})
	server.Use(cors.New())
	return &Server{
		Server: server,
	}
}
