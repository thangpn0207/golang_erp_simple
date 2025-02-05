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

// Create Purchase Order
func (c *PurchaseOrderController) CreatePurchaseOrder(ctx *gin.Context) {
	var request CreatePurchaseOrderRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order := models.PurchaseOrder{
		SupplierID:  request.SupplierID,
		UserID:      request.UserID,
		TotalAmount: request.TotalAmount,
		Status:      request.Status,
	}

	config.DB.Create(&order)
	ctx.JSON(http.StatusCreated, order)
}

// Get All Purchase Orders
func (c *PurchaseOrderController) GetPurchaseOrders(ctx *gin.Context) {
	var orders []models.PurchaseOrder
	config.DB.Preload("Supplier").Preload("User").Find(&orders)
	ctx.JSON(http.StatusOK, orders)
}
