package env

import (
	"os"

	"github.com/joho/godotenv"
)

const (
	DEBUG   = "debug"
	PROD    = "prod"
	GO_MODE = "mode"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		panic(err) //there is no .env file
	}
	mode := os.Getenv(GO_MODE)
	if mode == "" {
		mode = DEBUG
		os.Setenv(GO_MODE, mode)
	}
}
