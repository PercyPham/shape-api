package rest

import (
	"log"
	"net/http"
	"shape-api/internal/common/config"
	"shape-api/internal/repo"
	"strconv"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

func NewServer(cfg *Config) *server {
	return &server{
		r:        gin.Default(),
		userRepo: cfg.UserRepo,
	}
}

type Config struct {
	UserRepo repo.User
}

type server struct {
	r *gin.Engine

	userRepo repo.User
}

func (s *server) GetRouter() *gin.Engine {
	return s.r
}

func (s *server) Run() {
	switch config.App().ENV {
	case config.DEV:
		s.runDev()
	case config.PROD:
		s.runProd()
	default:
		panic("unsupported environment")
	}
}

func (s *server) runDev() {
	s.r.Use(CORSMiddleware())
	s.setupAPIs()

	port := strconv.Itoa(config.App().Port)
	s.r.Run(":" + port)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Methods", "*")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func (s *server) runProd() {
	s.setupAPIs()

	s.r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	log.Fatal(autotls.Run(s.r, config.App().Domains...))
}
