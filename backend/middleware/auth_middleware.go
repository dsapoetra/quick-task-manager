package middleware

import (
	"backend/pkg/jwt"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(401).JSON(fiber.Map{"error": "Authorization header required"})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(401).JSON(fiber.Map{"error": "Invalid authorization header format"})
		}

		token := parts[1]

		claims, err := jwt.ValidateToken(token, "your-secret-key")
		if err != nil {
			fmt.Printf("Token validation error: %v\n", err)
			return c.Status(401).JSON(fiber.Map{"error": "Invalid token"})
		}

		// Make sure we're accessing the correct claim key
		userID, ok := claims["user_id"]
		if !ok {
			return c.Status(401).JSON(fiber.Map{"error": "User ID not found in token"})
		}

		// Set the user ID in context
		c.Locals("userId", int64(userID.(float64)))

		return c.Next()
	}
}
