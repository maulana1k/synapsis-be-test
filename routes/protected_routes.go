package routes

import (
	"synapsis-be-test/handlers"
	"synapsis-be-test/middlewares"

	"github.com/gofiber/fiber/v2"
)

// ProtectedRoutes registers protected routes
func ProtectedRoutes(app *fiber.App) {
	// Create a group for protected routes with JWT middleware
	routes := app.Group("/api", middlewares.JwtMiddleware)

	// Register route to get products by category
	routes.Get("/products", handlers.GetByCategory)

	// Register route to add a new product
	routes.Post("/product/add", handlers.AddProduct)

	// Register route to get products in the cart
	routes.Get("/carts", handlers.GetProductInCart)

	// Register route to add a product to the cart
	routes.Post("/cart/add", handlers.AddProductToCart)

	// Register route to remove a product from the cart
	routes.Post("/cart/remove", handlers.RemoveProductFromCart)

	// Register route to checkout
	routes.Post("/cart/checkout", handlers.Checkout)
}
