package models

// Inventory Model
type Inventory struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	ProductID uint    `gorm:"not null" json:"product_id"`
	Quantity  int     `gorm:"default:0" json:"quantity"`
	Product   Product `gorm:"foreignKey:ProductID" json:"product"`
}
