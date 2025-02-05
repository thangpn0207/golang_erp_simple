// main.go
package main

import (
	"erp-be/config"
	"erp-be/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// Initialize configurations
	config.Init()

	// Set up router
	r := gin.Default()

	// CORS configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Setup routes
	routes.SetupRoutes(r)

	// Start server
	log.Fatal(r.Run(":8080"))
}
