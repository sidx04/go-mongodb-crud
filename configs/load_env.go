package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvironmentVariables() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Couldn't load environment variables...")
		return ""
	}

	log.Println("Loaded environment variables...")

	return os.Getenv("MONGO_URI")
}
