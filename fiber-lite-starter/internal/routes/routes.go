package routes

import (
	"fiber-lite-starter/config"
	"fiber-lite-starter/internal/repository/port"
	"fiber-lite-starter/pkg/response"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"strings"
)

type RouteRegistry struct {
	Repository port.RepositoryRegistry
	Validator  *validator.Validate
}

func NewRouteRegistry(repository port.RepositoryRegistry) *RouteRegistry {
	validate := validator.New()
	return &RouteRegistry{
		Repository: repository,
		Validator:  validate,
	}
}

func (r *RouteRegistry) RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Middleware X-API-KEY
	api.Use(func(c *fiber.Ctx) error {
		key := c.Get("x-api-key")
		if !strings.EqualFold(key, config.Envs.APIKeys.XApiKey) {
			log.Error().Msg("route::RegisterRoutes - Invalid x-api-key")
			return c.Status(fiber.StatusUnauthorized).JSON(response.Error("Unauthorized: invalid x-api-key"))
		}
		return c.Next()
	})

	// Public routes
	user := api.Group("/user")
	RegisterUserRoutes(user, r.Repository)

	// Fallback route: not found
	api.All("/*", func(c *fiber.Ctx) error {
		log.Info().
			Str("url", c.OriginalURL()).
			Str("method", c.Method()).
			Str("path", c.Path()).
			Str("query", c.Context().QueryArgs().String()).
			Str("ua", c.Get("User-Agent")).
			Str("ip", c.IP()).
			Msg("Route not found.")

		return c.Status(fiber.StatusNotFound).JSON(response.Error("Route not found"))
	})
}
