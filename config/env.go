package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func LoadEnv() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
