package main

import (
	"shape-api/internal/adapter/http/rest"
	"shape-api/internal/common/config"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env") // load `.env` if has
	godotenv.Load(".default.env")
	config.Load()

	s := rest.NewServer()
	s.Run()
}
