package controllers

import (
	"erp-be/config"
	"erp-be/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PurchaseOrderController struct{}

type CreatePurchaseOrderRequest struct {
	SupplierID  uint    `json:"supplier_id" binding:"required"`
	UserID      uint    `json:"user_id" binding:"required"`
	TotalAmount float64 `json:"total_amount"`
	Status      string  `json:"status"`
}
type CreatePurchaseOrderResponse struct {
	UserID     uint `json:"user_id"`
	SupplierID uint `json:"supplier_id"`
	Products   []struct {
		ProductID uint    `json:"product_id"`
		Price     float64 `json:"price"`
		Quantity  int     `json:"quantity"`
	} `json:"products"`
}

// Create Purchase Order
func (ctrl *PurchaseOrderController) CreatePurchaseOrder(c *gin.Context) {
	var request CreatePurchaseOrderResponse

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate product IDs
	for _, item := range request.Products {
		var product models.Product
		if err := config.DB.First(&product, item.ProductID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found"})
			return
		}
	}

	// Calculate total amount
	totalAmount := 0.0
	for _, item := range request.Products {
		totalAmount += item.Price * float64(item.Quantity)
	}

	// Start transaction
	tx := config.DB.Begin()

	purchaseOrder := models.PurchaseOrder{
		UserID:      request.UserID,
		SupplierID:  request.SupplierID,
		TotalAmount: totalAmount,
	}

	// Save Purchase Order
	if err := tx.Create(&purchaseOrder).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Process each product in the purchase order
	for _, item := range request.Products {
		purchaseOrderItem := models.PurchaseOrderItem{
			PurchaseOrderID: purchaseOrder.ID,
			ProductID:       item.ProductID,
			Price:           item.Price,
			Quantity:        item.Quantity,
		}
		if err := tx.Create(&purchaseOrderItem).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Update Inventory
		var inventory models.Inventory
		if err := tx.Where("product_id = ?", item.ProductID).First(&inventory).Error; err != nil {
			inventory = models.Inventory{
				ProductID: item.ProductID,
				Quantity:  item.Quantity,
			}
			tx.Create(&inventory)
		} else {
			inventory.Quantity += item.Quantity
			tx.Save(&inventory)
		}
	}

	// Commit transaction
	tx.Commit()

	c.JSON(http.StatusCreated, purchaseOrder)
}

// Get All Purchase Orders
func (ctrl *PurchaseOrderController) GetPurchaseOrders(c *gin.Context) {
	var orders []models.PurchaseOrder
	config.DB.Find(&orders)
	c.JSON(http.StatusOK, orders)
}
