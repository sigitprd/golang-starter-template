package routes

import (
	"echo-lite-starter/internal/handler"
	"echo-lite-starter/internal/repository/port"
	"echo-lite-starter/internal/service"
	"echo-lite-starter/pkg/response"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"net/http"
)

func RegisterUserRoutes(g *echo.Group, repo port.RepositoryRegistry) {
	userService := service.NewUserService(repo)
	userHandler := handler.NewUserHandler(userService)

	g.POST("", userHandler.CreateUser)
	g.GET("", userHandler.GetUsers)
	g.GET("/:id", userHandler.GetUserById)

	g.Any("/*", func(c echo.Context) error {
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
