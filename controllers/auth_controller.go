package controllers

import (
	"erp-be/config"
	"erp-be/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct{}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

func (auth *AuthController) Login(c *gin.Context) {
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", request.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials 1"})
		return
	}

	//if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
	//	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials 2"})
	//	return
	//}

	//token, err := utils.GenerateJWT(user.ID, user.Email, user.Role)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
	//	return
	//}

	c.JSON(http.StatusOK, gin.H{"token": "token"})
}

// Register User
func (c *AuthController) RegisterUser(ctx *gin.Context) {
	var request RegisterRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	//if err != nil {
	//	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
	//	return
	//}

	user := models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		Role:     "Admin",
	}
	config.DB.Create(&user)
	ctx.JSON(http.StatusCreated, gin.H{"token": "token"})
}
