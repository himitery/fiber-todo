package jwt

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v3"
)

var (
	// ErrJWTMissingOrMalformed is returned when the JWT is missing or malformed.
	ErrJWTMissingOrMalformed = errors.New("can't find jwt access token")
)

type jwtExtractor func(ctx fiber.Ctx) (string, error)

// jwtFromHeader returns a function that extracts token from the request header.
func jwtFromHeader(header string, authScheme string) func(ctx fiber.Ctx) (string, error) {
	return func(ctx fiber.Ctx) (string, error) {
		auth := ctx.Get(header)
		l := len(authScheme)
		if len(auth) > l+1 && strings.EqualFold(auth[:l], authScheme) {
			return strings.TrimSpace(auth[l:]), nil
		}
		return "", ErrJWTMissingOrMalformed
	}
}

// jwtFromQuery returns a function that extracts token from the query string.
func jwtFromQuery(param string) func(ctx fiber.Ctx) (string, error) {
	return func(ctx fiber.Ctx) (string, error) {
		token := ctx.Query(param)
		if token == "" {
			return "", ErrJWTMissingOrMalformed
		}
		return token, nil
	}
}

// jwtFromParam returns a function that extracts token from the url param string.
func jwtFromParam(param string) func(ctx fiber.Ctx) (string, error) {
	return func(ctx fiber.Ctx) (string, error) {
		token := ctx.Params(param)
		if token == "" {
			return "", ErrJWTMissingOrMalformed
		}
		return token, nil
	}
}

// jwtFromCookie returns a function that extracts token from the named cookie.
func jwtFromCookie(name string) func(ctx fiber.Ctx) (string, error) {
	return func(ctx fiber.Ctx) (string, error) {
		token := ctx.Cookies(name)
		if token == "" {
			return "", ErrJWTMissingOrMalformed
		}
		return token, nil
	}
}
