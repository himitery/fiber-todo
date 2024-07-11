package port

import "github.com/himitery/fiber-todo/internal/core/domain"

type JwtUsecase interface {
	Generate(auth domain.Auth) (domain.Token, error)
	Parse(token string) (domain.Claims, error)
}
