package models

import "gorm.io/gorm"

// Sales Order Model
type SalesOrder struct {
	gorm.Model
	ID          uint     `gorm:"primaryKey"`
	CustomerID  uint     `gorm:"not null"`
	UserID      uint     `gorm:"not null"`
	TotalAmount float64  `gorm:"default:0.00"`
	Status      string   `gorm:"type:enum('Pending','Completed','Cancelled');default:'Pending'"`
	Customer    Customer `gorm:"foreignKey:CustomerID"`
	User        User     `gorm:"foreignKey:UserID"`
}

// Sales Order Item Model
type SalesOrderItem struct {
	gorm.Model

	ID           uint       `gorm:"primaryKey"`
	SalesOrderID uint       `gorm:"not null"`
	ProductID    uint       `gorm:"not null"`
	Quantity     int        `gorm:"not null"`
	Price        float64    `gorm:"not null"`
	SalesOrder   SalesOrder `gorm:"foreignKey:SalesOrderID"`
	Product      Product    `gorm:"foreignKey:ProductID"`
}
