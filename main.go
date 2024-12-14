package main

import (
	_ "main/docs" // Swagger docs

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @Summary Get User
// @Description Get user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param   id     path    int     true  "User ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /users/{id} [get]
func getUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"user_id": id})
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server for demonstrating Swagger with Golang.
// @host localhost:8080
// @BasePath /
func main() {
	r := gin.Default()

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API routes
	r.GET("/users/:id", getUser)

	r.Run(":8080")

}
