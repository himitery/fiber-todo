package application

import (
	"fmt"

	"github.com/himitery/fiber-todo/internal/core/domain"
	"github.com/himitery/fiber-todo/internal/core/port"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	authRepository port.AuthRepository
}

func NewAuthService(authRepository port.AuthRepository) port.AuthUsecase {
	return &AuthService{
		authRepository: authRepository,
	}
}

func (service AuthService) SignIn(req port.SignInReq) (domain.Auth, error) {
	auth, err := service.authRepository.FindByEmail(req.Email)
	if err != nil {
		return domain.Auth{}, &port.PortError{Code: 400, Message: "사용자를 찾을 수 없습니다."}
	}

	if !service.validatePassword(auth.Password, req.Password) {
		return domain.Auth{}, &port.PortError{Code: 400, Message: "아이디 혹은 비밀번호가 틀렸습니다."}
	}

	return auth, nil
}

func (service AuthService) SignUp(req port.SignUpReq) (domain.Auth, error) {
	_, err := service.authRepository.FindByEmail(req.Email)
	if err == nil {
		return domain.Auth{}, &port.PortError{Code: 400, Message: fmt.Sprintf("이메일이 이미 존재합니다. (email: %s)", req.Email)}
	}

	password, err := service.generateHashedPassword(req.Password)
	if err != nil {
		return domain.Auth{}, &port.PortError{Code: 500, Message: err.Error()}
	}

	auth, err := service.authRepository.Save(&domain.Auth{
		Email:    req.Email,
		Password: password,
		Username: req.Username,
	})
	if err != nil {
		return domain.Auth{}, &port.PortError{Code: 500, Message: err.Error()}
	}

	return auth, nil
}

func (service AuthService) ReIssue(req port.ReIssueReq) (domain.Auth, error) {
	auth, err := service.authRepository.FindById(req.Id)
	if err != nil {
		return domain.Auth{}, &port.PortError{Code: 400, Message: "사용자를 찾을 수 없습니다."}
	}

	return auth, nil
}

func (service AuthService) UpdatePassword(req port.UpdatePasswordReq) (domain.Auth, error) {
	auth, err := service.authRepository.FindById(req.Id)
	if err != nil {
		return domain.Auth{}, &port.PortError{Code: 400, Message: "사용자를 찾을 수 없습니다."}
	}
	if !service.validatePassword(auth.Password, req.AlreadyPassword) {
		return domain.Auth{}, &port.PortError{Code: 400, Message: "비밀번호가 틀렸습니다."}
	}

	newPassword, err := service.generateHashedPassword(req.NewPassword)
	if err != nil {
		return domain.Auth{}, &port.PortError{Code: 500, Message: err.Error()}
	}

	auth.Password = newPassword
	service.authRepository.UpdatePassword(&auth)

	return auth, nil
}

func (service AuthService) generateHashedPassword(password string) (string, error) {
	res, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	return string(res), err
}

func (service AuthService) validatePassword(hashedPassword string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}
