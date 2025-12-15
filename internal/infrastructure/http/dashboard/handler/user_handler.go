package handler

import (
	"net/http"
	"strconv"

	"github.com/futek/donation-campaign/internal/application"
	"github.com/futek/donation-campaign/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log" // New import
)

// UserHandler handles user-related HTTP requests.
type UserHandler struct {
	userService application.UserService
}

// NewUserHandler creates a new UserHandler.
func NewUserHandler(userService application.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// CreateUser handles the creation of a new user.
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Error().Err(err).Msg("Failed to bind user JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.userService.CreateUser(c.Request.Context(), &user); err != nil {
		log.Error().Err(err).Msg("Failed to create user")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// GetUserByID handles retrieving a user by their ID.
func (h *UserHandler) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error().Err(err).Msg("Invalid user ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	user, err := h.userService.GetUserByID(c.Request.Context(), id)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get user by ID")
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
