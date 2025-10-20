package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load()
	
	if err != nil {
		log.Println("Warning: no .env file found, falling back to system envs")
	}
}