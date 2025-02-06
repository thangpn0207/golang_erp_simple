package config

import (
	"erp-be/models"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var DB *gorm.DB

func Init() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Correct MySQL DSN format
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)
	fmt.Println("DSN:", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}
	DB = db

	// Auto migrate models
	migrateModels()
}

func migrateModels() {
	DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Customer{},
		&models.Supplier{},
		&models.SalesOrder{},
		&models.SalesOrderItem{},
		&models.PurchaseOrder{},
		&models.PurchaseOrderItem{},
		&models.Inventory{},
	)
}
