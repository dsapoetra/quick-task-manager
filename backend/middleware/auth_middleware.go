package middleware

import (
	"github.com/gofiber/fiber/v2"
	"strings"
	"backend/pkg/jwt"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(401).JSON(fiber.Map{"error": "Authorization header required"})
		}

		// Check if the header starts with "Bearer "
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(401).JSON(fiber.Map{"error": "Invalid authorization header format"})
		}

		token := parts[1]
		
		// Validate token
		claims, err := jwt.ValidateToken(token, "your-secret-key") // TODO: Use config for secret
		if err != nil {
			return c.Status(401).JSON(fiber.Map{"error": "Invalid token"})
		}

		// Set user ID in context
		userId := claims["user_id"].(float64)
		c.Locals("userId", int64(userId))

		return c.Next()
	}
}
