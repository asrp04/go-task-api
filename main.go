package main

import (
	"go-task-api/config"
	"go-task-api/controllers"
	"go-task-api/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()
	router := gin.Default()
	router.POST("/api/signup", controllers.Signup)
	router.POST("/api/login", controllers.Login)
	protected := router.Group("/api")
	protected.Use(middleware.AuthRequired())
	{
		protected.POST("/tasks", controllers.CreateTask)
		protected.GET("/tasks", controllers.FindTasks)
		protected.PUT("/tasks/:id", controllers.UpdateTask)
		protected.DELETE("/tasks/:id", controllers.DeleteTask)
	}
	router.Run(":8080")
}
