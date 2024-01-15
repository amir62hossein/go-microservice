package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strings"
)

func Authentication() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(map[string]string{"message": "Authorization header is missing"})
		}

		// Split the Authorization header to get the token part
		authParts := strings.Fields(authHeader)
		if len(authParts) != 2 || authParts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(map[string]string{"message": "Invalid Authorization header format"})
		}

		token := authParts[1]

		// Parse the token
		parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("invalid signing method")
			}
			return []byte(os.Getenv("KEY")), nil
		})

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(map[string]string{"error": "Invalid token via signing method"})
		}

		if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
			userID, ok := claims["id"].(string)
			if !ok {
				return c.Status(fiber.StatusUnauthorized).JSON(map[string]string{"error": "Invalid token when to extract claim id"})
			}

			userEmail, ok := claims["email"].(string)
			if !ok {
				return c.Status(fiber.StatusUnauthorized).JSON(map[string]string{"error": "Invalid token when to extract claims email"})
			}

			// Set user ID and email in locals
			c.Locals("userID", userID)
			c.Locals("userEmail", userEmail)

			return c.Next()
		}

		return c.Status(fiber.StatusUnauthorized).JSON(map[string]string{"error": "Invalid token"})
	}
}
