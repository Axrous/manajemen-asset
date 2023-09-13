package helper

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// generate bcrypt
func HashPassword(password string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", fmt.Errorf("Failed to hash password %v", err.Error())
	}

	return string(hashedPass), nil
}

// compare password
func ComparePassword(hashedPassword, password string) error {
	hashedPass := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return hashedPass
}

func HashPasswordForgotPass(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil //convert to string
}
