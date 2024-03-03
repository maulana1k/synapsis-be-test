package routes

import (
	"synapsis-be-test/handlers"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(app *fiber.App) {
	route := app.Group("/api")

	route.Post("/auth/register", handlers.Register)
	route.Post("/auth/login", handlers.Login)

}
