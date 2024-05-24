package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"Assignment_3_Defense/models"
	"Assignment_3_Defense/utils"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	logger := utils.NewLogger()
	userHandler := NewUserHandler(logger)

	router := gin.Default()
	router.POST("/users", userHandler.CreateUser)

	t.Run("Create User Success", func(t *testing.T) {
		body := `{"name": "John Doe", "email": "john@example.com"}`
		req, _ := http.NewRequest("POST", "/users", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Contains(t, w.Body.String(), `"name":"John Doe"`)
		assert.Contains(t, w.Body.String(), `"email":"john@example.com"`)
	})

	t.Run("Create User Invalid Request", func(t *testing.T) {
		body := `{"name": "John Doe"}`
		req, _ := http.NewRequest("POST", "/users", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), `"error":"Invalid request"`)
	})
}

func TestGetUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	logger := utils.NewLogger()
	userHandler := NewUserHandler(logger)

	router := gin.Default()
	router.GET("/users/:id", userHandler.GetUser)

	t.Run("Get User Success", func(t *testing.T) {
		service := userHandler.service
		service.CreateUser(&models.User{Name: "John Doe", Email: "john@example.com"})

		req, _ := http.NewRequest("GET", "/users/1", nil)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), `"name":"John Doe"`)
		assert.Contains(t, w.Body.String(), `"email":"john@example.com"`)
	})

	t.Run("Get User Not Found", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/users/999", nil)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Contains(t, w.Body.String(), `"error":"User not found"`)
	})
}

func TestDeleteUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	logger := utils.NewLogger()
	userHandler := NewUserHandler(logger)

	router := gin.Default()
	router.DELETE("/users/:id", userHandler.DeleteUser)

	t.Run("Delete User Success", func(t *testing.T) {
		service := userHandler.service
		service.CreateUser(&models.User{Name: "John Doe", Email: "john@example.com"})

		req, _ := http.NewRequest("DELETE", "/users/1", nil)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), `"message":"User deleted"`)
	})

	t.Run("Delete User Not Found", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/users/999", nil)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Contains(t, w.Body.String(), `"error":"User not found"`)
	})
}
