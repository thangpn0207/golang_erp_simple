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
func (c *SalesOrderController) CreateSalesOrder(ctx *gin.Context) {
	var request CreateSalesOrderRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order := models.SalesOrder{
		CustomerID:  request.CustomerID,
		UserID:      request.UserID,
		TotalAmount: request.TotalAmount,
		Status:      request.Status,
	}

	config.DB.Create(&order)
	ctx.JSON(http.StatusCreated, order)
}

// Get All Sales Orders
func (c *SalesOrderController) GetSalesOrders(ctx *gin.Context) {
	var orders []models.SalesOrder
	config.DB.Preload("Customer").Preload("User").Find(&orders)
	ctx.JSON(http.StatusOK, orders)
}
