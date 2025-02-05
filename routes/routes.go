package routes

import (
	"erp-be/controllers"
	"erp-be/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{})
	})
	// Auth routes
	auth := controllers.AuthController{}
	r.POST("/api/auth/login", auth.Login)
	r.POST("/api/auth/register", auth.RegisterUser)

	// Protected routes
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		// Products
		products := controllers.ProductController{}
		api.GET("/products", products.GetProducts)
		api.POST("/products", products.CreateProduct)
		//api.PUT("/products/:id", products.Update)
		//api.DELETE("/products/:id", products.Delete)

		// Sales Orders
		sales := controllers.SalesOrderController{}
		api.GET("/sales-orders", sales.GetSalesOrders)
		api.POST("/sales-orders", sales.CreateSalesOrder)
		//api.GET("/sales-orders/:id", sales.Get)
		//api.PUT("/sales-orders/:id", sales.Update)

		// Purchase Orders
		purchases := controllers.PurchaseOrderController{}
		api.GET("/purchase-orders", purchases.GetPurchaseOrders)
		api.POST("/purchase-orders", purchases.CreatePurchaseOrder)
		//api.GET("/purchase-orders/:id", purchases.Get)
		//api.PUT("/purchase-orders/:id", purchases.Update)
		//

		// Customer
		customer := controllers.CustomerController{}
		api.GET("/customers", customer.GetCustomers)
		api.POST("/customers", customer.GetCustomers)
		//api.GET("/purchase-orders/:id", purchases.Get)
		//api.PUT("/purchase-orders/:id", purchases.Update)
		//

		// Purchase Orders
		supplier := controllers.SupplierController{}
		api.GET("/suppliers", supplier.GetSuppliers)
		api.POST("/suppliers", supplier.CreateSupplier)
		//api.GET("/purchase-orders/:id", purchases.Get)
		//api.PUT("/purchase-orders/:id", purchases.Update)
		//
	}
}
