package main

import (
	"shape-api/internal/adapter/http/rest"
	"shape-api/internal/adapter/repoimpl/mysqlrepo"
	"shape-api/internal/common/config"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load() // load `.env` if has
	config.Load()

	db, err := mysqlrepo.Connect(config.MySQL().DSN)
	if err != nil {
		panic(err)
	}

	server, err := rest.NewServer(&rest.Config{
		UserRepo:      mysqlrepo.NewUserRepo(db),
		TriangleRepo:  mysqlrepo.NewTriangleRepo(db),
		RectangleRepo: mysqlrepo.NewRectangleRepo(db),
		SquareRepo:    mysqlrepo.NewSquareRepo(db),
	})
	if err != nil {
		panic(err)
	}

	server.Run()
}
