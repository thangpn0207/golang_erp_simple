package controllers

import (
	"erp-be/config"
	"erp-be/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SupplierController struct{}

// Create Supplier
func (c *SupplierController) CreateSupplier(ctx *gin.Context) {
	var request struct {
		Name    string `json:"name" binding:"required"`
		Contact string `json:"contact"`
		Address string `json:"address"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	supplier := models.Supplier{
		Name:    request.Name,
		Contact: request.Contact,
		Address: request.Address,
	}
	config.DB.Create(&supplier)
	ctx.JSON(http.StatusCreated, supplier)
}

// Get All Suppliers
func (c *SupplierController) GetSuppliers(ctx *gin.Context) {
	var suppliers []models.Supplier
	config.DB.Find(&suppliers)
	ctx.JSON(http.StatusOK, suppliers)
}
