package controllers

import (
	"erp-be/config"
	"erp-be/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SalesOrderController struct{}
type CreateSalesOrderRequest struct {
	CustomerID  uint    `json:"customer_id" binding:"required"`
	UserID      uint    `json:"user_id" binding:"required"`
	TotalAmount float64 `json:"total_amount"`
	Status      string  `json:"status"`
}

// Create Sales Order
func (ctrl *SalesOrderController) CreateSalesOrder(c *gin.Context) {
	var request struct {
		UserID     uint `json:"user_id"`
		CustomerID uint `json:"customer_id"`
		Products   []struct {
			ProductID uint    `json:"product_id"`
			Price     float64 `json:"price"`
			Quantity  int     `json:"quantity"`
		} `json:"products"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate product IDs and check inventory stock
	for _, item := range request.Products {
		var product models.Product
		if err := config.DB.First(&product, item.ProductID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found"})
			return
		}

		var inventory models.Inventory
		if err := config.DB.Where("product_id = ?", item.ProductID).First(&inventory).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found in inventory"})
			return
		}

		if inventory.Quantity < item.Quantity {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Not enough stock for product " + product.Name})
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

	salesOrder := models.SalesOrder{
		UserID:      request.UserID,
		CustomerID:  request.CustomerID,
		TotalAmount: totalAmount,
	}

	// Save Sales Order
	if err := tx.Create(&salesOrder).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Process each product in the sales order
	for _, item := range request.Products {
		salesOrderItem := models.SalesOrderItem{
			SalesOrderID: salesOrder.ID,
			ProductID:    item.ProductID,
			Price:        item.Price,
			Quantity:     item.Quantity,
		}
		if err := tx.Create(&salesOrderItem).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Update Inventory
		var inventory models.Inventory
		if err := tx.Where("product_id = ?", item.ProductID).First(&inventory).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Inventory not found"})
			return
		}

		inventory.Quantity -= item.Quantity
		if err := tx.Save(&inventory).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	// Commit transaction
	tx.Commit()

	c.JSON(http.StatusCreated, salesOrder)
}

// Get All Sales Orders
// controllers/sales_order_controller.go
func (ctrl *SalesOrderController) GetSalesOrders(c *gin.Context) {
	var orders []models.SalesOrder

	// Preload items associated with each sales order
	if err := config.DB.Preload("Items").Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}
