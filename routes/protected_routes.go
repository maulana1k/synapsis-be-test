package routes

import (
	"synapsis-be-test/handlers"
	"synapsis-be-test/middlewares"

	"github.com/gofiber/fiber/v2"
)

func ProtectedRoutes(app *fiber.App) {
	routes := app.Group("/api", middlewares.JwtMiddleware)

	routes.Get("/products", handlers.GetByCategory)
	routes.Post("/product/add", handlers.AddProduct)
	routes.Get("/carts", handlers.GetProductInCart)
	routes.Post("/cart/add", handlers.AddProductToCart)
	routes.Post("/cart/remove", handlers.RemoveProductFromCart)
	routes.Post("/cart/checkout", handlers.Checkout)
}
