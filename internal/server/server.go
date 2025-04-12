package server

import (
	"github.com/labstack/echo/v4"
	"github.com/nest-hackathon/server/config"
	"github.com/nest-hackathon/server/internal/api/router"
)

type Server struct {
	config *config.Config
	app    *echo.Echo
}

func NewServer() *Server {
	app := echo.New()
	cfg := config.LoadConfig()

	// Register routes
	router.RegisterRoutes(app)

	return &Server{
		config: cfg,
		app:    app,
	}
}

func (s *Server) Start() error {
	return s.app.Start(":" + s.config.Port)
}
