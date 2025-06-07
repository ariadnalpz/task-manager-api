package middleware

import (
	"github.com/gofiber/fiber/v2"
	"task-manager-api/utils"
)

// AuthMiddleware verifica el token JWT
func AuthMiddleware(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token no proporcionado"})
	}

	// Eliminar "Bearer " del token si está presente
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	userID, err := utils.ValidateJWT(tokenString)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token inválido"})
	}

	// Guardar el ID del usuario en el contexto
	c.Locals("userID", userID)
	return c.Next()
}