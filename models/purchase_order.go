package models

import "gorm.io/gorm"

// Purchase Order Model
type PurchaseOrder struct {
	gorm.Model
	ID          uint    `gorm:"primaryKey" json:"id"`
	SupplierID  uint    `gorm:"not null" json:"supplier_id"`
	UserID      uint    `gorm:"not null" json:"user_id"`
	TotalAmount float64 `gorm:"default:0.00" json:"total_amount"`
	Status      string  `gorm:"type:enum('Pending','Completed','Cancelled');default:'Pending'"  json:"status"`
	User        User    `gorm:"foreignKey:UserID" json:"user"`
}

// Purchase Order Item Model
type PurchaseOrderItem struct {
	ID              uint          `gorm:"primaryKey" json:"id"`
	PurchaseOrderID uint          `gorm:"not null" json:"purchase_order_id"`
	ProductID       uint          `gorm:"not null" json:"product_id"`
	Quantity        int           `gorm:"not null" json:"quantity"`
	Price           float64       `gorm:"not null" json:"price"`
	PurchaseOrder   PurchaseOrder `gorm:"foreignKey:PurchaseOrderID" json:"purchase_order"`
	Product         Product       `gorm:"foreignKey:ProductID" json:"product"`
}
