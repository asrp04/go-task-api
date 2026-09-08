package main

import (
	"go-task-api/config"
	"go-task-api/controllers"
	"go-task-api/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. डेटाबेस कनेक्शन शुरू करें
	config.ConnectDatabase()

	// 2. Gin राउटर सेटअप करें
	router := gin.Default()

	// 1. पब्लिक रूट्स (Public Routes - कोई भी एक्सेस कर सकता है)
	router.POST("/api/signup", controllers.Signup)
	router.POST("/api/login", controllers.Login)

	// 2. सुरक्षित रूट्स ग्रुप (Protected Routes Group)
	// इसके अंदर जो भी रूट होगा, वो पहले AuthRequired मिडिलवेयर से गुजरेगा
	protected := router.Group("/api")
	protected.Use(middleware.AuthRequired())
	{
		protected.POST("/tasks", controllers.CreateTask)
		protected.GET("/tasks", controllers.FindTasks)
		protected.PUT("/tasks/:id", controllers.UpdateTask)
		protected.DELETE("/tasks/:id", controllers.DeleteTask)
	}

	// 4. सर्वर रन करें
	router.Run(":8080")
}
