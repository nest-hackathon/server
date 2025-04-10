package server

import (

	"github.com/labstack/echo/v4"
	"github.com/nest-hackathon/server/config"
)

type Server struct {
	config *config.Config
	app    *echo.Echo
	router *echo.Group
}

func NewServer() *Server {
	app := echo.New()
	cfg := config.LoadConfig()

	return &Server{
		config: cfg,
		app:    app,
		router: app.Group("/api/v1"),
	}
}

func (s *Server) Start() error {
	return s.app.Start(":" + s.config.Port)
}