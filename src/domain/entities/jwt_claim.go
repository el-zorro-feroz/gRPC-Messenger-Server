package entities

import (
	"github.com/golang-jwt/jwt/v4"
)

type JwtClaim struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	jwt.StandardClaims
}

type JwtRefreshClaim struct {
	ID string `json:"id"`
	jwt.StandardClaims
}
