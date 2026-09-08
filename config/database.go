package config

import (
	"fmt"
	"go-task-api/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// .env फाइल को लोड कर रहे हैं
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: .env file not found, using system env")
	}

	// os.Getenv के जरिए .env से वैल्यूज निकालना
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// डायनामिक DSN स्ट्रिंग बनाना
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("डेटाबेस से कनेक्ट करने में विफल!")
	}

	database.AutoMigrate(&models.Task{}, &models.User{})

	fmt.Println("डेटाबेस कनेक्शन सफल रहा (.env लोड हो गया)!")
	DB = database
}
