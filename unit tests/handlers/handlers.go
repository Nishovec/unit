package handlers

import (
	"net/http"
	"strconv"

	"Assignment_3_Defense/models"
	"Assignment_3_Defense/services"
	"Assignment_3_Defense/utils"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *services.UserService
	logger  *utils.Logger
}

func NewUserHandler(logger *utils.Logger) *UserHandler {
	return &UserHandler{
		service: services.NewUserService(),
		logger:  logger,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		h.logger.Error("Failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	h.service.CreateUser(&user)
	h.logger.Info("User created: %v", user)
	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error("Invalid user ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, found := h.service.GetUser(id)
	if !found {
		h.logger.Warn("User not found: %d", id)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	h.logger.Info("User retrieved: %v", user)
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error("Invalid user ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		h.logger.Error("Failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := h.service.UpdateUser(id, &user); err != nil {
		h.logger.Warn("User not found: %d", id)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	h.logger.Info("User updated: %v", user)
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error("Invalid user ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := h.service.DeleteUser(id); err != nil {
		h.logger.Warn("User not found: %d", id)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	h.logger.Info("User deleted: %d", id)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
