package router

import (
	"github.com/gofiber/fiber/v3"
	server "github.com/himitery/fiber-todo/internal"
)

type HealthRouter struct {
	Router fiber.Router
}

func NewHealthRouter(httpServer *server.HttpServer) {
	router := HealthRouter{
		Router: httpServer.Server.Group("/api/health"),
	}

	router.Init()
}

func (router HealthRouter) Init() {
	router.Router.Get("/", router.Ping)
}

func (router HealthRouter) Ping(ctx fiber.Ctx) error {
	return ctx.Status(200).SendString("ok")
}
