package controllers

import (
	"erp-be/config"
	"erp-be/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CustomerController struct{}

// Create Customer
func (c *CustomerController) CreateCustomer(ctx *gin.Context) {
	var request struct {
		Name    string `json:"name" binding:"required"`
		Contact string `json:"contact"`
		Address string `json:"address"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer := models.Customer{
		Name:    request.Name,
		Contact: request.Contact,
		Address: request.Address,
	}
	config.DB.Create(&customer)
	ctx.JSON(http.StatusCreated, customer)
}

// Get All Customers
func (c *CustomerController) GetCustomers(ctx *gin.Context) {
	var customers []models.Customer
	config.DB.Find(&customers)
	ctx.JSON(http.StatusOK, customers)
}
