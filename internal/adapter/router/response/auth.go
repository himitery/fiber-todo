package response

import "github.com/himitery/fiber-todo/internal/core/domain"

type TokenRes struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func NewTokenRes(token domain.Token) *TokenRes {
	return &TokenRes{
		AccessToken:  token.Access,
		RefreshToken: token.Refresh,
	}
}
