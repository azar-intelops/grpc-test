package configs

import (
	"os"

	"github.com/joho/godotenv"
)

func GetMongoURL() string {
	err := godotenv.Load(".env.local")
	if err != nil {
		return err.Error()
	}
	return os.Getenv("MONGODB_URL")
}
