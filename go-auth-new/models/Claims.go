package model

import "github.com/golang-jwt/jwt"

type Claims struct {
	Role string `json:"role"`
	jwt.StandardClaims
}
