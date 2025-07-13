package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"net/http"
)

func AuthRole(authorizedRoles []string) echo.MiddlewareFunc {
	// convert slice to map for O(1) lookup
	allowed := make(map[string]struct{}, len(authorizedRoles))
	for _, r := range authorizedRoles {
		allowed[r] = struct{}{}
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			role, ok := c.Get("role").(string)
			if !ok {
				return c.JSON(http.StatusForbidden, map[string]any{
					"message": "Terlarang: role tidak ditemukan",
					"success": false,
				})
			}

			if _, found := allowed[role]; found {
				return next(c)
			}

			log.Warn().Any("payload", map[string]any{
				"role":             role,
				"authorized_roles": authorizedRoles,
			}).Msg("middleware::AuthRole - Unauthorized")

			return c.JSON(http.StatusForbidden, map[string]any{
				"message": "Terlarang: role anda tidak diizinkan untuk mengakses resource ini",
				"success": false,
			})
		}
	}
}
