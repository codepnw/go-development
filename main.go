package main

import (
	"log"

	"github.com/codepnw/godevelopment/internal/api"
	"github.com/joho/godotenv"
)

const (
	version = "v1"
	envFile = "test.env"
)

func main() {
	if err := godotenv.Load(envFile); err != nil {
		log.Fatal("cannot load .env file")
	}

	api.NewRoutes(version)
}
