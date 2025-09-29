package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func Admin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		isAdmin, ok := c.Locals("is_admin").(bool)
		if !ok || !isAdmin {
			return c.Status(403).JSON(fiber.Map{"error": "Admin only"})
		}
		return c.Next()
	}
}
