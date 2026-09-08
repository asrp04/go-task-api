package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"-" gorm:"not null"` // json:"-" का मतलब है कि रिस्पॉन्स में पासवर्ड कभी बाहर नहीं जाएगा
}