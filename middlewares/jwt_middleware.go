package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// JwtMiddleware is a middleware for JWT authentication
func JwtMiddleware(c *fiber.Ctx) error {
	// Extract JWT token from the Authorization header
	authHeader := c.Get("Authorization")

	// Check if the Authorization header is missing
	if authHeader == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "Missing token")
	}

	// Check if the Authorization header has the correct format
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid token format")
	}

	// Extract the token string without the "Bearer " prefix
	tokenString := authHeader[len("Bearer "):]

	// Parse and validate the JWT token
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Validate token signing method here if needed
		return []byte("synapsis"), nil
	})

	// Check for token parsing errors
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
	}

	// Check if the token is valid
	if token.Valid {
		// Set user ID from claims in the context for further use in handlers
		userID := claims["ID"].(string)
		c.Locals("user", userID)
		return c.Next()
	}

	// Token is invalid
	return fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
}
