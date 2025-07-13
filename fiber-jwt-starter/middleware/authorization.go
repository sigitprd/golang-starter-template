package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func AuthRole(authorizedRoles []string) func(*fiber.Ctx) error {
	// Buat map untuk lookup cepat (O(1))
	roleSet := make(map[string]struct{}, len(authorizedRoles))
	for _, role := range authorizedRoles {
		roleSet[role] = struct{}{}
	}

	return func(c *fiber.Ctx) error {
		role, ok := c.Locals("role").(string)
		if !ok {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "Terlarang: role tidak ditemukan dalam context",
				"success": false,
			})
		}

		if _, exists := roleSet[role]; exists {
			return c.Next()
		}

		log.Warn().
			Str("role", role).
			Strs("authorized_roles", authorizedRoles).
			Msg("middleware::AuthRole - Unauthorized role access attempt")

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Terlarang: role anda tidak diizinkan untuk mengakses resource ini",
			"success": false,
		})
	}
}
