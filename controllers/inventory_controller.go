package controllers

import (
	"erp-be/config"
	"erp-be/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type InventoryController struct{}

func (ctrl *InventoryController) GetInventory(c *gin.Context) {
	var inventory []models.Inventory
	config.DB.Find(&inventory)
	c.JSON(http.StatusOK, inventory)
}
