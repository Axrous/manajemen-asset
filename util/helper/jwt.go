package helper

import (
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

var secret = []byte(os.Getenv("JWT_SECRET"))

type JWTClaims struct {
	UserID string `json:"email"`
	jwt.StandardClaims
}

// init jwt in here
func GenerateJWT(Email string) (string, error) {
	claims := JWTClaims{
		UserID: Email,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(6 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

// parse jwt
func ParseJWT(tokenHeader string) (*jwt.Token, error) {
	return jwt.Parse(tokenHeader, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
}
