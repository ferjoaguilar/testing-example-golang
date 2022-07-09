package utils

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func Enviroment(v string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", err
	}

	if v == "" {
		return "", errors.New("Enviroment variable is required")
	}

	varEnv := os.Getenv(v)
	return varEnv, nil
}
