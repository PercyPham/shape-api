package rest

import (
	"log"
	"net/http"
	"shape-api/internal/common/config"
	"strconv"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

func NewServer() *server {
	return &server{
		r: gin.Default(),
	}
}

type server struct {
	r *gin.Engine
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
