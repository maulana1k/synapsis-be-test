package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JwtMiddleware(c *fiber.Ctx) error {
	// Extract JWT token from the Authorization header
	authHeader := c.Get("Authorization")

	// Check if the Authorization header is empty
	if authHeader == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "Missing token")
	}

	// Check if the Authorization header has the Bearer prefix
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid token format")
	}

	// Extract the token string without the "Bearer " prefix
	tokenString := authHeader[len("Bearer "):]

	// Parse the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check token signing method

		return []byte("synapsis"), nil
	})

	// Check for parsing errors
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
	}

	// Check if the token is valid
	if token.Valid {
		// Set user claims in the context for further use in handlers
		claims := token.Claims.(jwt.MapClaims)
		c.Locals("user", claims["ID"])
		return c.Next()
	} else {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
	}
}
