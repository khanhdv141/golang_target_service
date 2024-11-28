package util

import (
	"CMS/config"
	"crypto/rsa"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

const RS256 = "RS256"

type JWTUtils interface {
	GenerateToken(jwt.Claims) (string, error)
	ParseToken(string) (jwt.Claims, error)
}

type jwtUtils struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func NewJWTUtils() JWTUtils {
	pubKeyFileContent, err := os.ReadFile(config.ApplicationConfig.JWT.PublicKeyFilePath)
	if err != nil {
		panic(err)
	}
	privateKeyFileContent, err := os.ReadFile(config.ApplicationConfig.JWT.PrivateKeyFilePath)
	if err != nil {
		panic(err)
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(pubKeyFileContent)
	if err != nil {
		panic(err)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyFileContent)
	return &jwtUtils{
		privateKey: privateKey,
		publicKey:  publicKey,
	}
}

func (jwtUtils *jwtUtils) GenerateToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(jwtUtils.privateKey)
}

func (jwtUtils *jwtUtils) ParseToken(tokenString string) (jwt.Claims, error) {
	var claims jwt.MapClaims
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return jwtUtils.publicKey, nil
	})
	if err != nil {
		return nil, err
	}
	return claims, nil
}
