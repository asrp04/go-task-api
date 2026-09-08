package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model        // यह ID, CreatedAt, UpdatedAt, DeletedAt अपने आप देगा
	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description"`
	Done        bool   `json:"done" gorm:"default:false"`
}