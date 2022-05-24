package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Common = loadCommonConfigurations()

func GetEnvVar(key string) string {
	appEnv := os.Getenv("APP_ENV")
	var err interface{}
	if appEnv == "prod" {
		err = godotenv.Load("configs/prod.env")
	} else if appEnv == "staging" {
		err = godotenv.Load("configs/staging.env")
	} else if appEnv == "dev" {
		err = godotenv.Load("configs/dev.env")
	} else {
		err = godotenv.Load("configs/local.env")
	}
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	return os.Getenv(key)
}
