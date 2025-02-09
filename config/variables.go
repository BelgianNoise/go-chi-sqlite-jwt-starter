package config

import (
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type ServerVariables struct {
	Port             string `env:"PORT" env-default:"9494"`
	DB_FOLDER        string `env:"DB_FOLDER" env-default:"./data"`
	AUTH_PRIVATE_KEY string `env:"AUTH_PRIVATE_KEY" env-required:"true"`
}

var Variables ServerVariables

func InitializeVariables() {
	log.Printf("Loading environment variables...")
	defer log.Printf("Environment variables loaded!")

	err := godotenv.Load()
	if err != nil {
		errorMsg, _ := fmt.Printf("Error loading .env file: %v", err)
		panic(errorMsg)
	}
	err = cleanenv.ReadEnv(&Variables)

	if err != nil {
		panic(err)
	}
}
