package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID          uint    `gorm:"primaryKey"`
	Name        string  `gorm:"size:255;not null"`
	Description string  `gorm:"type:text"`
	Price       float64 `gorm:"not null"`
	StockLevel  int     `gorm:"default:0"`
}
