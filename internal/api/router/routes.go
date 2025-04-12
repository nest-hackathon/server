package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nest-hackathon/server/internal/api/handlers"
)

func RegisterRoutes(e *echo.Echo) {
	// Health check endpoint
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "ok",
		})
	})

	api := e.Group("/api/v1")

	

	robotHandler := handlers.NewRobotHandler(nil)
	api.POST("/robot", robotHandler.CreateRobot)
	api.DELETE("/robot/:id", robotHandler.DeleteRobot)


	authHandler := handlers.NewAuthHandler(nil)
	api.POST("/login", authHandler.Login)
	api.POST("/register", authHandler.Register)

}
