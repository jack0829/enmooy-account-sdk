package jwt

import (
	"github.com/golang-jwt/jwt/v4"
)

type LoginClaims struct {
	jwt.RegisteredClaims
	Data *Data `json:"data"`
}
