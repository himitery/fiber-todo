package domain

import gojwt "github.com/golang-jwt/jwt/v5"

type Token struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

type Claims struct {
	gojwt.Claims `json:"-"`
	Sub          string `json:"sub"`
	Exp          int64  `json:"exp"`
}
