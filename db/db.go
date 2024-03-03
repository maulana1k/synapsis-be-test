package db

import (
	"synapsis-be-test/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// Open the SQLite database file using Gorm
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("Error opening database: " + err.Error())
	}

	// Auto-migrate the necessary models
	if err := db.AutoMigrate(&models.User{}, &models.Product{}, &models.Cart{}, &models.CartItems{}, &models.Transaction{}); err != nil {
		panic("Error migrating database: " + err.Error())
	}

	// Assign the database connection to the global variable
	DB = db
}
