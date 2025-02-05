package models

import "gorm.io/gorm"

// Inventory Model
type Inventory struct {
	gorm.Model
	ID        uint    `gorm:"primaryKey"`
	ProductID uint    `gorm:"not null"`
	Quantity  int     `gorm:"default:0"`
	Product   Product `gorm:"foreignKey:ProductID"`
}
