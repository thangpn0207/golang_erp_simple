package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"size:100;not null"`
	Email    string `gorm:"size:100;unique;not null"`
	Role     string `gorm:"type:enum('Admin','Sales','Purchase','Inventory');not null"`
	Password string `json:"password"`
}
