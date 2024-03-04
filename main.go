package main

import (
	"synapsis-be-test/db"
	"synapsis-be-test/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Create a new Fiber app
	app := fiber.New()

	// Initialize the database
	db.InitDB()

	// Use logger middleware
	app.Use(logger.New())

	// Register public routes
	routes.PublicRoutes(app)

	// Register protected routes
	routes.ProtectedRoutes(app)

	// Start the server on port 3000
	app.Listen(":3000")
}
