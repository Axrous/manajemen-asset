package helper

import (
	"fmt"
	"github.com/joho/godotenv"
)

// load env
func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading .env file %v ", err.Error())
	}
	return nil
}
