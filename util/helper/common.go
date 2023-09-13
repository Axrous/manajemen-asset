package helper

import (
	"github.com/google/uuid"
	"regexp"
	"unicode"
)

func GenerateUUID() string {
	return uuid.NewString()
}

func ContainsUppercase(s string) bool {
	for _, char := range s {
		if unicode.IsUpper(char) {
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
