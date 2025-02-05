package models

import "gorm.io/gorm"

// Purchase Order Model
type PurchaseOrder struct {
	gorm.Model
	ID          uint     `gorm:"primaryKey"`
	SupplierID  uint     `gorm:"not null"`
	UserID      uint     `gorm:"not null"`
	TotalAmount float64  `gorm:"default:0.00"`
	Status      string   `gorm:"type:enum('Pending','Completed','Cancelled');default:'Pending'"`
	Supplier    Supplier `gorm:"foreignKey:SupplierID"`
	User        User     `gorm:"foreignKey:UserID"`
}

// Purchase Order Item Model
type PurchaseOrderItem struct {
	gorm.Model
	ID              uint          `gorm:"primaryKey"`
	PurchaseOrderID uint          `gorm:"not null"`
	ProductID       uint          `gorm:"not null"`
	Quantity        int           `gorm:"not null"`
	Price           float64       `gorm:"not null"`
	PurchaseOrder   PurchaseOrder `gorm:"foreignKey:PurchaseOrderID"`
	Product         Product       `gorm:"foreignKey:ProductID"`
}
