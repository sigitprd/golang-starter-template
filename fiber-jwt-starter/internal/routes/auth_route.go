package routes

import (
	"fiber-jwt-starter/internal/handler"
	"fiber-jwt-starter/internal/repository/port"
	"fiber-jwt-starter/internal/service"
	"fiber-jwt-starter/middleware"
	"fiber-jwt-starter/pkg/response"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func RegisterAuthRoutes(router fiber.Router, repo port.RepositoryRegistry) {
	authService := service.NewAuthService(repo)
	authHandler := handler.NewAuthHandler(authService)

	router.Post("/login", authHandler.Login)
	router.Post("/refresh", authHandler.Refresh)
	router.Post("/register", authHandler.Register)

	// Protected route
	protected := router.Group("/me", middleware.AuthBearer)
	protected.Get("/", authHandler.Profile)

	// Catch-all for unknown routes under /auth
	router.All("/*", func(c *fiber.Ctx) error {
		log.Info().
			Str("url", c.OriginalURL()).
			Str("method", c.Method()).
			Str("path", c.Path()).
			Str("query", string(c.Context().QueryArgs().QueryString())).
			Str("ua", c.Get("User-Agent")).
			Str("ip", c.IP()).
			Msg("Route not found.")
		return c.Status(fiber.StatusNotFound).JSON(response.Error("Route not found"))
	})
}
