package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func GetEnv(env string) []string {

	godotenv.Load("config.env")

	configuration := strings.Split(os.Getenv(env), ",")

	return configuration
}
