package middleware

import (
	"fiber-jwt-starter/pkg/jwthandler"
	"github.com/gofiber/fiber/v2"
	"strings"

	"github.com/rs/zerolog/log"
)

// AuthBearer adalah middleware untuk validasi JWT Bearer token
func AuthBearer(c *fiber.Ctx) error {
	accessToken := c.Get("Authorization")
	unauthorizedResponse := fiber.Map{
		"message": "Unauthorized",
		"success": false,
	}

	if accessToken == "" || !strings.HasPrefix(accessToken, "Bearer ") {
		log.Error().Msg("middleware::AuthBearer - Unauthorized [Missing or invalid Authorization header]")
		return c.Status(fiber.StatusUnauthorized).JSON(unauthorizedResponse)
	}

	// remove the "Bearer " prefix
	tokenString := strings.TrimPrefix(accessToken, "Bearer ")

	claims, err := jwthandler.ParseToken(tokenString)
	if err != nil {
		log.Error().
			Err(err).
			Str("token", tokenString).
			Msg("middleware::AuthBearer - Error while parsing token")
		return c.Status(fiber.StatusUnauthorized).JSON(unauthorizedResponse)
	}

	c.Locals("user_id", claims.ID)
	c.Locals("role", claims.Role)

	return c.Next()
}

func GetUserIDFromContext(c *fiber.Ctx) string {
	id, _ := c.Locals("user_id").(string)
	return id
}

func GetRoleFromContext(c *fiber.Ctx) string {
	role, _ := c.Locals("role").(string)
	return role
}
