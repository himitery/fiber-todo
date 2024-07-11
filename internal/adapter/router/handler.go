package router

import (
	"github.com/gofiber/fiber/v3"
	"github.com/himitery/fiber-todo/config"
	"github.com/himitery/fiber-todo/config/jwt"
)

func JwtHandler(conf *config.Config) func(fiber.Ctx) error {
	return jwt.New(jwt.Config{
		SigningKey: jwt.SigningKey{Key: []byte(conf.Jwt.AccessKey)},
		ContextKey: "auth",
		AuthScheme: "Bearer",
	})
}
