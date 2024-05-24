package main

import (
	"Assignment_3_Defense/handlers"
	"Assignment_3_Defense/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	logger := utils.NewLogger()
	userHandler := handlers.NewUserHandler(logger)

	router.POST("/users", userHandler.CreateUser)
	router.GET("/users/:id", userHandler.GetUser)
	router.PUT("/users/:id", userHandler.UpdateUser)
	router.DELETE("/users/:id", userHandler.DeleteUser)

	logger.Info("Starting server on port 8080")
	router.Run(":8080")
}
