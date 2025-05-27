// pkg/config/config.go
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Load() string {
	if err := godotenv.Load("configs/.env"); err != nil {
		log.Println("No .env file, using env vars")
	}
	return os.Getenv("RPC_URL")
}
