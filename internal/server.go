package server

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/himitery/fiber-todo/config"
	"github.com/himitery/fiber-todo/config/validator"
	"github.com/himitery/fiber-todo/internal/adapter/router/response"
	"github.com/himitery/fiber-todo/internal/core/port"
)

type HttpServer struct {
	Server *fiber.App
}

func NewHttpServer(conf *config.Config) *HttpServer {
	server := fiber.New(fiber.Config{
		ErrorHandler:    errorHandler,
		StructValidator: validator.NewStructValidator(),
	})

	server.Use(cors.New(cors.Config{
		AllowOrigins:     conf.Cors.Origins,
		AllowMethods:     conf.Cors.Methods,
		AllowHeaders:     conf.Cors.Headers,
		AllowCredentials: conf.Cors.Credentials,
	}))

	server.Use(logger.New())
	server.Use(recover.New())

	return &HttpServer{
		Server: server,
	}
}

func (httpServer HttpServer) Listen(address string) error {
	return httpServer.Server.Listen(address, fiber.ListenConfig{
		EnablePrintRoutes: true,
	})
}

func (httpServer HttpServer) Shutdown() error {
	return httpServer.Server.Shutdown()
}

func errorHandler(ctx fiber.Ctx, err error) error {
	switch err := err.(type) {
	case *port.PortError:
		return ctx.Status(err.Code).JSON(response.ErrorRes{Message: err.Message})
	default:
		return ctx.Status(500).JSON(response.ErrorRes{Message: err.Error()})
	}
}
