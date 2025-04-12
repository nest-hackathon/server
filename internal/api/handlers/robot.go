package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nest-hackathon/server/internal/api/models"
	"github.com/nest-hackathon/server/internal/database"
)

type RobotHandler struct {
	db *database.DatabaseImpl
}

func NewRobotHandler(db *database.DatabaseImpl) *RobotHandler {
	return &RobotHandler{db: db}
}

func (h *RobotHandler) CreateRobot(c echo.Context) error {
	var robot models.Robot
	if err := c.Bind(&robot); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	// Validate required fields
	if robot.Name == "" || robot.Type == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Name and type are required"})
	}

	robot.Status = "inactive"
	robot.Battery = 100

	return c.JSON(http.StatusCreated, robot)
}

func (h *RobotHandler) DeleteRobot(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Robot ID is required"})
	}


	return c.JSON(http.StatusOK, map[string]string{"message": "Robot deleted successfully"})
}
