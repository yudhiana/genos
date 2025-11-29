package db

import (
	"log"
	"sync"

	"genos/modules/orders/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbOnce sync.Once
var dbConnection *gorm.DB

// GetDatabaseConnection returns the database connection with a singleton pattern
func GetDatabaseConnection() *gorm.DB {
	dbOnce.Do(func() {
		// connect to database
		// This will automatically create 'orders.db' if it doesn't exist
		db, err := gorm.Open(sqlite.Open("orders.db"), &gorm.Config{})
		if err != nil {
			log.Fatal("failed to connect database:", err)
		}

		dbConnection = db.Debug()

		// Create the table (if not exists)
		err = dbConnection.AutoMigrate(&models.Order{}, &models.OrderItem{})
		if err != nil {
			log.Fatal("failed to migrate schema:", err)
		}

		log.Println("SQLite database and orders table created successfully!")
	})
	return dbConnection
}
