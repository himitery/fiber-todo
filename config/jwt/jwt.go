package jwt

import (
	"reflect"

	"github.com/gofiber/fiber/v3"
	gojwt "github.com/golang-jwt/jwt/v5"
)

var (
	defaultTokenLookup = "header:" + fiber.HeaderAuthorization
)

// New ...
func New(config ...Config) fiber.Handler {
	cfg := makeCfg(config)

	extractors := cfg.getExtractors()

	// Return middleware handler
	return func(ctx fiber.Ctx) error {
		// Filter request to skip middleware
		if cfg.Filter != nil && cfg.Filter(ctx) {
			return ctx.Next()
		}
		var auth string
		var err error

		for _, extractor := range extractors {
			auth, err = extractor(ctx)
			if auth != "" && err == nil {
				break
			}
		}
		if err != nil {
			return cfg.ErrorHandler(ctx, err)
		}
		var token *gojwt.Token

		if _, ok := cfg.Claims.(gojwt.MapClaims); ok {
			token, err = gojwt.Parse(auth, cfg.KeyFunc)
		} else {
			t := reflect.ValueOf(cfg.Claims).Type().Elem()
			claims := reflect.New(t).Interface().(gojwt.Claims)
			token, err = gojwt.ParseWithClaims(auth, claims, cfg.KeyFunc)
		}
		if err == nil && token.Valid {
			// Store user information from token into context.
			sub, _ := token.Claims.GetSubject()
			ctx.Locals(cfg.ContextKey, sub)
			return cfg.SuccessHandler(ctx)
		}
		return cfg.ErrorHandler(ctx, err)
	}
}
