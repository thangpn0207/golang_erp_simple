package middleware

import (
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//authHeader := c.GetHeader("Authorization")
		//if authHeader == "" {
		//	c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
		//	c.Abort()
		//	return
		//}
		//
		//bearerToken := strings.Split(authHeader, " ")
		//if len(bearerToken) != 2 {
		//	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
		//	c.Abort()
		//	return
		//}

		//claims, err := utils.ValidateJWT(bearerToken[1])
		//if err != nil {
		//	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		//	c.Abort()
		//	return
		//}
		//
		//c.Set("userID", claims.UserID)
		//c.Set("userRole", claims.Role)
		c.Next()
	}
}
