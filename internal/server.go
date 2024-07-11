package server

import (
	"fmt"

	goValidator "github.com/go-playground/validator/v10"
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
	case goValidator.ValidationErrors:
		for _, it := range err {
			switch it.Tag() {
			case "required":
				return ctx.Status(400).JSON(response.ErrorRes{Message: fmt.Sprintf("%s is required.", it.Field())})
			case "password":
				return ctx.Status(400).JSON(response.ErrorRes{Message: "비밀번호는 영어 대소문자 혹은 특수 문자로 시작하며, 각 최소 1개의 영어 대소문자, 특수 문자, 숫자를 포함해야 합니다. 또한 최소 8자리 이상을 만족해야합니다."})
			}
		}
		return ctx.Status(400).JSON(response.ErrorRes{Message: err.Error()})
	default:
		return ctx.Status(500).JSON(response.ErrorRes{Message: err.Error()})
	}
}
