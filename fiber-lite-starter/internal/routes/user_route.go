package routes

import (
	"fiber-lite-starter/internal/handler"
	"fiber-lite-starter/internal/repository/port"
	"fiber-lite-starter/internal/service"
	"fiber-lite-starter/pkg/response"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func RegisterUserRoutes(router fiber.Router, repo port.RepositoryRegistry) {
	userService := service.NewUserService(repo)
	userHandler := handler.NewUserHandler(userService)

	router.Post("", userHandler.CreateUser)
	router.Get("", userHandler.GetUsers)
	router.Get("/:id", userHandler.GetUserById)

	// Catch-all for unknown routes under /user
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
