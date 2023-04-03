package configs

import (
	"log"

	"github.com/joho/godotenv"
)

func InitEnv() {
	if err := godotenv.Load("configs/.env"); err != nil {
		log.Fatal("Error loading .env file")
	}
}
