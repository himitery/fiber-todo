package port

import (
	"github.com/google/uuid"
	"github.com/himitery/fiber-todo/internal/core/domain"
)

type AuthRepository interface {
	FindById(id uuid.UUID) (domain.Auth, error)
	FindByEmail(email string) (domain.Auth, error)
	Save(auth *domain.Auth) (domain.Auth, error)
	UpdatePassword(auth *domain.Auth) (domain.Auth, error)
}

type AuthUsecase interface {
	SignIn(req SignInReq) (domain.Auth, error)
	SignUp(req SignUpReq) (domain.Auth, error)
	ReIssue(req ReIssueReq) (domain.Auth, error)
	UpdatePassword(req UpdatePasswordReq) (domain.Auth, error)
}

type SignInReq struct {
	Email    string
	Password string
}

type SignUpReq struct {
	Email    string
	Password string
	Username string
}

type ReIssueReq struct {
	Id uuid.UUID
}

type UpdatePasswordReq struct {
	Id              uuid.UUID
	AlreadyPassword string
	NewPassword     string
}
