package application

import (
	"errors"
	"time"

	gojwt "github.com/golang-jwt/jwt/v5"
	"github.com/himitery/fiber-todo/config"
	"github.com/himitery/fiber-todo/internal/core/domain"
	"github.com/himitery/fiber-todo/internal/core/port"
)

type JwtService struct {
	Algorithm     *gojwt.SigningMethodHMAC
	AccessSecret  []byte
	RefreshSecret []byte
}

func NewJwtService(conf *config.Config) port.JwtUsecase {

	return &JwtService{
		Algorithm:     gojwt.SigningMethodHS256,
		AccessSecret:  []byte(conf.Jwt.AccessKey),
		RefreshSecret: []byte(conf.Jwt.RefreshKey),
	}
}

func (service JwtService) Generate(auth domain.Auth) (domain.Token, error) {
	accessToken, err := gojwt.NewWithClaims(service.Algorithm, domain.Claims{
		Sub: auth.Id.String(),
		Exp: time.Now().Add(time.Hour * 1).Unix(),
	}).SignedString(service.AccessSecret)
	if err != nil {
		return domain.Token{}, err
	}

	refreshToken, err := gojwt.NewWithClaims(service.Algorithm, domain.Claims{
		Sub: auth.Id.String(),
		Exp: time.Now().Add(time.Hour * 24 * 30).Unix(),
	}).SignedString(service.RefreshSecret)
	if err != nil {
		return domain.Token{}, err
	}

	return domain.Token{Access: accessToken, Refresh: refreshToken}, nil
}

func (service JwtService) Parse(token string) (domain.Claims, error) {
	res, err := gojwt.ParseWithClaims(token, gojwt.MapClaims{}, service.keyFunc(service.AccessSecret))
	if err == nil {
		sub, _ := res.Claims.GetSubject()
		return domain.Claims{Sub: sub}, nil
	}

	res, err = gojwt.ParseWithClaims(token, gojwt.MapClaims{}, service.keyFunc(service.RefreshSecret))
	if err == nil {
		sub, _ := res.Claims.GetSubject()
		return domain.Claims{Sub: sub}, nil

	}

	return domain.Claims{}, &port.PortError{Message: "유효하지 않은 토큰입니다."}
}

func (service JwtService) keyFunc(key []byte) func(*gojwt.Token) (interface{}, error) {
	return func(token *gojwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*gojwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return key, nil
	}
}
