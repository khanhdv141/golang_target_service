package dto

import "github.com/golang-jwt/jwt/v5"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Token struct {
	Token     string `json:"token"`
	TokenType string `json:"token_type"`
}

type JwtClaims struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
	UserId   uint   `json:"user_id"`
}
