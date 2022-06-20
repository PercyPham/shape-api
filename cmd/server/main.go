package main

import (
	"shape-api/internal/adapter/http/rest"
	"shape-api/internal/adapter/repoimpl/mysqlrepo"
	"shape-api/internal/common/config"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env") // load `.env` if has
	godotenv.Load(".default.env")
	config.Load()

	db, err := mysqlrepo.Connect(config.MySQL().DSN)
	if err != nil {
		panic(err)
	}

	s := rest.NewServer(&rest.Config{
		UserRepo: mysqlrepo.NewUserRepo(db),
	})

	s.Run()
}
