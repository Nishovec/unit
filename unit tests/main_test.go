package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"Assignment_3_Defense/handlers"
	"Assignment_3_Defense/utils"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	logger := utils.NewLogger()
	userHandler := handlers.NewUserHandler(logger)

	router.POST("/users", userHandler.CreateUser)
	router.GET("/users/:id", userHandler.GetUser)
	router.PUT("/users/:id", userHandler.UpdateUser)
	router.DELETE("/users/:id", userHandler.DeleteUser)

	return router
}

func TestCreateAndGetUserIntegration(t *testing.T) {
	router := setupRouter()

	// Create User
	t.Run("Create User", func(t *testing.T) {
		body := `{"name": "Jane Doe", "email": "jane@example.com"}`
		req, _ := http.NewRequest("POST", "/users", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Contains(t, w.Body.String(), `"name":"Jane Doe"`)
		assert.Contains(t, w.Body.String(), `"email":"jane@example.com"`)
	})

	// Get User
	t.Run("Get User", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/users/1", nil)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), `"name":"Jane Doe"`)
		assert.Contains(t, w.Body.String(), `"email":"jane@example.com"`)
	})
}

func TestUpdateAndDeleteUserIntegration(t *testing.T) {
	router := setupRouter()

	// Create User for updating and deleting
	body := `{"name": "John Smith", "email": "john.smith@example.com"}`
	reqCreate, _ := http.NewRequest("POST", "/users", strings.NewReader(body))
	reqCreate.Header.Set("Content-Type", "application/json")

	wCreate := httptest.NewRecorder()
	router.ServeHTTP(wCreate, reqCreate)
	assert.Equal(t, http.StatusCreated, wCreate.Code)

	// Update User
	t.Run("Update User", func(t *testing.T) {
		updateBody := `{"name": "John Smith Updated", "email": "john.updated@example.com"}`
		reqUpdate, _ := http.NewRequest("PUT", "/users/1", strings.NewReader(updateBody))
		reqUpdate.Header.Set("Content-Type", "application/json")

		wUpdate := httptest.NewRecorder()
		router.ServeHTTP(wUpdate, reqUpdate)

		assert.Equal(t, http.StatusOK, wUpdate.Code)
		assert.Contains(t, wUpdate.Body.String(), `"name":"John Smith Updated"`)
		assert.Contains(t, wUpdate.Body.String(), `"email":"john.updated@example.com"`)
	})

	// Delete User
	t.Run("Delete User", func(t *testing.T) {
		reqDelete, _ := http.NewRequest("DELETE", "/users/1", nil)

		wDelete := httptest.NewRecorder()
		router.ServeHTTP(wDelete, reqDelete)

		assert.Equal(t, http.StatusOK, wDelete.Code)
		assert.Contains(t, wDelete.Body.String(), `"message":"User deleted"`)
	})
}
