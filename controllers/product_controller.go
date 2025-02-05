package controllers

import (
	"erp-be/config"
	"erp-be/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductController struct{}
type CreateProductRequest struct {
	Name  string  `json:"name" binding:"required"`
	Price float64 `json:"price" binding:"required"`
}

// Create Product
func (c *ProductController) CreateProduct(ctx *gin.Context) {
	var request CreateProductRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := models.Product{
		Name:  request.Name,
		Price: request.Price,
	}
	config.DB.Create(&product)

	// Initialize inventory for the new product
	inventory := models.Inventory{
		ProductID: product.ID,
		Quantity:  0,
	}
	config.DB.Create(&inventory)

	ctx.JSON(http.StatusCreated, product)
}

// Get All Products
func (c *ProductController) GetProducts(ctx *gin.Context) {
	var products []models.Product
	config.DB.Find(&products)
	ctx.JSON(http.StatusOK, products)
}
