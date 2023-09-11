package helper

import (
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"unicode"
)

func ContainsUppercase(s string) bool {
	for _, char := range s {
		if unicode.IsUpper(char) {
			return true
		}
	}
	return false
}

func ContainsNumber(s string) bool {
	for _, char := range s {
		if unicode.IsNumber(char) {
			return true
		}
	}
	return false
}

func ContainsSpecialChar(s string) bool {
	// Regular expression to match any special character
	re := regexp.MustCompile(`[!@#$%^&*()_+=\[{\]};:'",<.>/?]`)
	return re.MatchString(s)
}

func GenerateUUID() string {
	return uuid.NewString()
}

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
