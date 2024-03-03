package main

import (
	"synapsis-be-test/db"
	"synapsis-be-test/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	db.InitDB()

	// app.Use(swagger.New())

	app.Use(logger.New())

	routes.PublicRoutes(app)
	routes.ProtectedRoutes(app)

	app.Listen(":3000")
}
