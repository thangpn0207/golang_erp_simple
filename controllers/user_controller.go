package controllers

import (
	"erp-be/config"
	"erp-be/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct{}

// Get All Users
func (c *UserController) GetUsers(ctx *gin.Context) {
	var users []models.User
	config.DB.Find(&users)
	ctx.JSON(http.StatusOK, users)
}
