package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tarkanaciksoz/api-todo-app/internal/model"
)

func Init(logger *log.Logger) model.Config {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		logger.Printf("You must declare APP_ENV before run")
		os.Exit(1)
	}

	err := godotenv.Load(".env" + "." + appEnv)
	if err != nil {
		logger.Printf("Error while Read .env file: %s\n", err.Error())
		os.Exit(1)
	}

	return model.Config{
		AppEnv:      appEnv,
		BindAddress: os.Getenv("BIND_ADDRESS"),
	}
}
