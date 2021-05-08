package token

import (
	"time"

	"github.com/anonychun/go-blog-api/internal/config"
	jwt "github.com/dgrijalva/jwt-go"
)

type TokenModel interface {
	GenerateClaims() jwt.MapClaims
}

func GenerateToken(model TokenModel) (string, error) {
	claims := model.GenerateClaims()
	claims["exp"] = time.Now().Add(config.Cfg().JwtTTL).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Cfg().JwtSecretKey))
}

func ExtractClaims(tokenString string) jwt.MapClaims {
	tokenParse, _ := jwt.Parse(tokenString, nil)
	return tokenParse.Claims.(jwt.MapClaims)
}
