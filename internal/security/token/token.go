package token

import (
	"time"

	"github.com/anonychun/go-blog-api/internal/config"
	jwt "github.com/dgrijalva/jwt-go"
)

type Generator interface {
	GenerateClaims() jwt.MapClaims
}

func GenerateToken(g Generator) (string, error) {
	claims := g.GenerateClaims()
	claims["exp"] = time.Now().Add(config.Cfg().JwtTTL).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Cfg().JwtSecretKey))
}
