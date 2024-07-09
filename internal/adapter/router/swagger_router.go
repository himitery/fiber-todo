package router

import (
	"github.com/gofiber/fiber/v3"
	"github.com/himitery/fiber-todo/config/oas"
	server "github.com/himitery/fiber-todo/internal"

	docs "github.com/himitery/fiber-todo/docs"
)

type SwaggerRouter struct {
	Router fiber.Router
}

func NewSwaggerRouter(httpServer *server.HttpServer) {
	swaggerRouter := SwaggerRouter{
		httpServer.Server.Group("/docs"),
	}

	swaggerRouter.Init()
}

func (router SwaggerRouter) Init() {
	docs.SwaggerInfo.Title = "Fiber Todo"
	docs.SwaggerInfo.Description = "Todo application using golang fiber"
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.Schemes = []string{"https"}

	router.Router.Get("/*", oas.New(oas.Config{
		URL:          "swagger.json",
		Layout:       "StandaloneLayout",
		ValidatorUrl: "localhost",
	}))
}
