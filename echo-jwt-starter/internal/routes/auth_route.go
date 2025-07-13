package routes

import (
	"echo-jwt-starter/internal/handler"
	"echo-jwt-starter/internal/repository/port"
	"echo-jwt-starter/internal/service"
	"echo-jwt-starter/middleware"
	"echo-jwt-starter/pkg/response"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"net/http"
)

func RegisterAuthRoutes(g *echo.Group, repo port.RepositoryRegistry) {
	authService := service.NewAuthService(repo)
	authHandler := handler.NewAuthHandler(authService)

	g.POST("/login", authHandler.Login)
	g.POST("/refresh", authHandler.Refresh)
	g.POST("/register", authHandler.Register)

	// Protected route
	protected := g.Group("/me")
	protected.Use(middleware.AuthBearer)
	protected.GET("", authHandler.Profile)

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
