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

		// Validate token
		claims, err := jwt.ValidateToken(token, "your-secret-key")
		if err != nil {
			return c.Status(401).JSON(fiber.Map{"error": "Invalid token"})
		}

		// Set both the claims and user ID in context
		userId := claims["user_id"].(float64)
		c.Locals("userId", int64(userId))

		// Convert user_id to string for the Subject claim
		c.Locals("Bearer", map[string]interface{}{
			"Subject": fmt.Sprintf("%d", int64(userId)),
		})

		return c.Next()
	}
}
