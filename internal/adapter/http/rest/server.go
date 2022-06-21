package rest

import (
	"fmt"
	"log"
	"net/http"
	"shape-api/internal/common/config"
	"shape-api/internal/repo"
	"strconv"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

func NewServer(cfg *Config) (*server, error) {
	if err := cfg.validate(); err != nil {
		return nil, err
	}
	return &server{
		r:             gin.Default(),
		userRepo:      cfg.UserRepo,
		triangleRepo:  cfg.TriangleRepo,
		rectangleRepo: cfg.RectangleRepo,
		squareRepo:    cfg.SquareRepo,
	}, nil
}

type Config struct {
	UserRepo      repo.User
	TriangleRepo  repo.Triangle
	RectangleRepo repo.Rectangle
	SquareRepo    repo.Square
}

func (c *Config) validate() error {
	if c.UserRepo == nil {
		return fmt.Errorf("user repo must not be nil")
	}
	if c.TriangleRepo == nil {
		return fmt.Errorf("triangle repo must not be nil")
	}
	if c.RectangleRepo == nil {
		return fmt.Errorf("rectangle repo must not be nil")
	}
	if c.SquareRepo == nil {
		return fmt.Errorf("square repo must not be nil")
	}
	return nil
}

type server struct {
	r *gin.Engine

	userRepo      repo.User
	triangleRepo  repo.Triangle
	rectangleRepo repo.Rectangle
	squareRepo    repo.Square
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
