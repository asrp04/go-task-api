package controllers

import (
	"go-task-api/config"
	"go-task-api/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

// 1. नया टास्क बनाने के लिए (POST /api/tasks)
func CreateTask(c *gin.Context) {
	var input struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// डेटाबेस मॉडल तैयार कर रहे हैं
	task := models.Task{
		Title:       input.Title,
		Description: input.Description,
		Done:        false, // यह models.Task की फील्ड से मैच कर रहा है
	}

	config.DB.Create(&task)

	c.JSON(http.StatusCreated, task)
}

// 2. सारे टास्क देखने के लिए (GET /api/tasks)
func FindTasks(c *gin.Context) {
	var tasks []models.Task

	config.DB.Find(&tasks)

	c.JSON(http.StatusOK, tasks)
}

// 3. किसी टास्क को अपडेट करने के लिए (PUT /api/tasks/:id)
func UpdateTask(c *gin.Context) {
	var task models.Task
	id := c.Param("id") // URL से ID निकालना (जैसे /api/tasks/1)

	// पहले चेक करें कि वो टास्क डेटाबेस में है भी या नहीं
	if err := config.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "टास्क नहीं मिला!"})
		return
	}

	// इनपुट डेटा वैलिडेट करें
	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Done        bool   `json:"done"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// GORM के जरिए डेटाबेस में अपडेट करना (जैसे MySQL में UPDATE...SET)
	config.DB.Model(&task).Updates(models.Task{
		Title:       input.Title,
		Description: input.Description,
		Done:        input.Done,
	})

	c.JSON(http.StatusOK, task)
}

// 4. किसी टास्क को डिलीट करने के लिए (DELETE /api/tasks/:id)
func DeleteTask(c *gin.Context) {
	var task models.Task
	id := c.Param("id")

	// पहले चेक करें कि टास्क मौजूद है
	if err := config.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "टास्क नहीं मिला!"})
		return
	}

	// GORM के जरिए डिलीट करना (SELECT * FROM tasks WHERE id = ...)
	config.DB.Delete(&task)

	c.JSON(http.StatusOK, gin.H{"message": "टास्क सफलतापूर्वक डिलीट कर दिया गया!"})
}
