package models

import "gorm.io/gorm"

// Supplier Model
type Supplier struct {
	gorm.Model
	ID      uint   `gorm:"primaryKey"`
	Name    string `gorm:"size:255;not null"`
	Contact string `gorm:"size:100"`
	Address string `gorm:"type:text"`
}
