package helper

import (
	"final-project-enigma-clean/model"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

var secret = []byte(os.Getenv("JWT_SECRET"))

// init jwt in here
func GenerateJWT(userLogin model.UserLoginRequest) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   userLogin.ID,
		"issued_at": time.Now().Unix(),
		"exp_at":    time.Now().Add(6 * time.Hour),
	})
	return token.SignedString(secret)
}

// parse jwt
func ParseJWT(tokenHeader string) (*jwt.Token, error) {
	return jwt.Parse(tokenHeader, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
}
