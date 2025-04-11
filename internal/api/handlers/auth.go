package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nest-hackathon/server/internal/database"
	"github.com/nest-hackathon/server/internal/types"
	"github.com/nest-hackathon/server/pkg/jwt"
)

type Handler struct {
	db *database.DatabaseImpl
}

func NewHandler(db *database.DatabaseImpl) *Handler {
	return &Handler{db: db}
}

func (h *Handler) Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "" || password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Username and password are required"})
	}

	user, err := h.db.GetUserByUsername(username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get user"})
	}

	if user.Password != password {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
	}

	token, err := jwt.NewJWTConfig("your-secret-key").EncodeJWT(user.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate token"})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}

func (h *Handler) Register(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "" || password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Username and password are required"})
	}

	user := &types.User{
		Username: username,
		Password: password,
	}

	err := h.db.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User created successfully"})
}
