package routes

import (
	"echo-lite-starter/config"
	"echo-lite-starter/internal/repository/port"
	"echo-lite-starter/pkg/response"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type RouteRegistry struct {
	Repository port.RepositoryRegistry
}

func NewRouteRegistry(repository port.RepositoryRegistry) *RouteRegistry {
	return &RouteRegistry{
		Repository: repository,
	}
}

func (r *RouteRegistry) RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api")

	api.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: "header:x-api-key",
		Validator: func(key string, c echo.Context) (bool, error) {
			return strings.EqualFold(key, config.Envs.APIKeys.XApiKey), nil
		},
		ErrorHandler: func(err error, c echo.Context) error {
			log.Error().Err(err).Msg("route::SetupRoutes - Invalid x-api-key")
			return c.JSON(http.StatusUnauthorized, response.Error(err.Error()))
		},
	}))

	// Public routes
	user := api.Group("/user")
	RegisterUserRoutes(user, r.Repository)

	// Fallback route for handling unknown routes
	api.Any("/*", func(c echo.Context) error {
		log.Info().
			Str("url", c.Request().URL.String()). // c.Request().URL.String() is getting the full URL
			Str("method", c.Request().Method).    // c.Request().Method is getting the request method
			Str("path", c.Request().URL.Path).    // c.Request().URL.Path is getting the request path
			Str("query", c.QueryString()).        // c.QueryString() is getting the query string
			Str("ua", c.Request().UserAgent()).   // c.Request().UserAgent() is getting the user agent
			Str("ip", c.RealIP()).                // c.RealIP() is getting the real IP address
			Msg("Route not found.")
		return c.JSON(http.StatusNotFound, response.Error("Route not found"))
	})
}
