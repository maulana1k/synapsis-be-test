package routes

import (
	"synapsis-be-test/handlers"

	"github.com/gofiber/fiber/v2"
)

// PublicRoutes registers public routes
func PublicRoutes(app *fiber.App) {
	route := app.Group("/api")

	// Register route for user registration
	route.Post("/auth/register", handlers.Register)

	// Register route for user login
	route.Post("/auth/login", handlers.Login)
}
