package database

import (
	"time"

	"gorm.io/gorm"
)

// Model base struct for database models
type Model struct {
	Id        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type User struct {
	Model
	UserName string `gorm:"unique" json:"username"`
}
