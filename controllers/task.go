package controllers

import (
	"go-task-api/config"
	"go-task-api/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var input struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := models.Task{
		Title:       input.Title,
		Description: input.Description,
		Done:        false, 
	}

	config.DB.Create(&task)

	c.JSON(http.StatusCreated, task)
}

func FindTasks(c *gin.Context) {
	var tasks []models.Task

	config.DB.Find(&tasks)

	c.JSON(http.StatusOK, tasks)
}

func UpdateTask(c *gin.Context) {
	var task models.Task
	id := c.Param("id") 

	if err := config.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found!"})
		return
	}

	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Done        bool   `json:"done"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&task).Updates(models.Task{
		Title:       input.Title,
		Description: input.Description,
		Done:        input.Done,
	})

	c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
	var task models.Task
	id := c.Param("id")

	if err := config.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found!"})
		return
	}

	config.DB.Delete(&task)

	c.JSON(http.StatusOK, gin.H{"message": "Task removed successfully!"})
}
