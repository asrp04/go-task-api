package controllers

import (
	"net/http"
	"time"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go-task-api/config"
	"go-task-api/models"
	"golang.org/x/crypto/bcrypt"
)

// 1. यूजर साइनअप (Register)
func Signup(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// पासवर्ड को हैश (Encrypt) करना
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "पासवर्ड हैश करने में विफल"})
		return
	}

	user := models.User{
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	// डेटाबेस में सेव करना
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ईमेल पहले से मौजूद है!"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "रजिस्ट्रेशन सफल रहा!"})
}

// 2. यूजर लॉगिन (Login & Generate Token)
func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	// डेटाबेस में यूजर को ढूंढना
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "गलत ईमेल या पासवर्ड"})
		return
	}

	// पासवर्ड मैच करना
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "गलत ईमेल या पासवर्ड"})
		return
	}

	// JWT टोकन बनाना (जो 24 घंटे के लिए वैलिड होगा)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "टोकन बनाने में विफल"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
