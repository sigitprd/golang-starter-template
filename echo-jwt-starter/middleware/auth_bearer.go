package middleware

import (
	"echo-jwt-starter/pkg/jwthandler"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

// AuthBearer adalah middleware untuk validasi JWT Bearer token
func AuthBearer(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			log.Warn().Msg("middleware::AuthBearer - missing or invalid Authorization header")
			return c.JSON(http.StatusUnauthorized, map[string]any{
				"message": "Unauthorized",
				"success": false,
			})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := jwthandler.ParseToken(tokenString)
		if err != nil {
			log.Error().
				Err(err).
				Str("token", tokenString).
				Msg("middleware::AuthBearer - failed to parse token")
			return c.JSON(http.StatusUnauthorized, map[string]any{
				"message": "Unauthorized",
				"success": false,
			})
		}

		log.Debug().
			Str("user_id", claims.ID).
			Str("role", claims.Role).
			Str("subject", claims.Subject).
			Msg("middleware::AuthBearer - token validated")

		c.Set("user_id", claims.ID)
		c.Set("role", claims.Role)

		return next(c)
	}
}

// GetUserIDFromContext mengambil user ID dari context Echo
func GetUserIDFromContext(c echo.Context) string {
	id, _ := c.Get("user_id").(string)
	return id
}

// GetRoleFromContext mengambil user role dari context Echo
func GetRoleFromContext(c echo.Context) string {
	role, _ := c.Get("role").(string)
	return role
}
