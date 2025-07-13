package main

import (
	"echo-jwt-starter/config"
	"echo-jwt-starter/internal/repository/psql"
	"echo-jwt-starter/internal/routes"
	dbconfig "echo-jwt-starter/pkg/db"
	"echo-jwt-starter/pkg/logging"
	echovalidator "echo-jwt-starter/pkg/validator"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/time/rate"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	// Load .env
	config.LoadEnvs()

	// Setup logger
	logLevel, err := zerolog.ParseLevel(config.Envs.App.LogLevel)
	if err != nil {
		logLevel = zerolog.InfoLevel
	}
	logging.SetupLogger(config.Envs.App.Environment, config.Envs.App.LogFile, logLevel)

	// Init DB
	db, err := dbconfig.NewPostgresConnection()
	if err != nil {
		log.Fatal().Err(err).Msg("main:: failed to connect to database")
	}
	defer db.Close()

	// Echo instance
	e := echo.New()
	// Application Middlewares
	if !config.Envs.App.Environment.IsProd() {
		//app.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(50)))
		e.Use(middleware.RateLimiterWithConfig(middleware.RateLimiterConfig{
			Skipper: middleware.DefaultSkipper,
			Store: middleware.NewRateLimiterMemoryStoreWithConfig(
				middleware.RateLimiterMemoryStoreConfig{Rate: rate.Limit(50), Burst: 30, ExpiresIn: 30 * time.Second},
			),
			IdentifierExtractor: func(ctx echo.Context) (string, error) {
				id := ctx.RealIP()
				return id, nil
			},
			ErrorHandler: func(context echo.Context, err error) error {
				return context.JSON(http.StatusForbidden, nil)
			},
			DenyHandler: func(context echo.Context, identifier string, err error) error {
				return context.JSON(http.StatusTooManyRequests, nil)
			},
		}))
	}
	e.Use(middleware.CORS())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		//AllowMethods: "GET,POST,PUT,DELETE,PATCH,OPTIONS,HEAD",
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS", "HEAD"},
		//AllowHeaders: "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin,Authorization,aplication/json; charset=utf-8,x-api-key",
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Content-Length", "Accept-Language", "Accept-Encoding", "Connection", "Access-Control-Allow-Origin", "Authorization", "aplication/json; charset=utf-8", "x-api-key"},
		//AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderContentLength, echo.HeaderAcceptLanguage, echo.HeaderAcceptEncoding, echo.HeaderConnection, echo.HeaderAccessControlAllowOrigin, echo.HeaderAuthorization},
	}))
	e.Use(middleware.Gzip())
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		LogErrorFunc: func(c echo.Context, err error, stack []byte) error {
			log.Error().Err(err).Bytes("stack", stack).Msg("Panic occurred")
			return nil
		},
	}))
	e.Use(middleware.RequestID())
	//e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
	//	XSSProtection:         "1; mode=block",
	//	ContentTypeNosniff:    "nosniff",
	//	XFrameOptions:         "DENY",
	//	HSTSMaxAge:            3600,
	//	HSTSExcludeSubdomains: true,
	//	HSTSPreloadEnabled:    false,
	//	ContentSecurityPolicy: "default-src 'self'",
	//	ReferrerPolicy:        "no-referrer",
	//}))
	e.Use(middleware.Secure())
	e.Validator = echovalidator.NewValidator() // Set custom validator

	// Route registry
	routeRegistry := routes.NewRouteRegistry(
		psql.NewRepositoryRegistry(db),
	)
	routeRegistry.RegisterRoutes(e)

	// Start server
	// Run server in goroutine
	go func() {
		serverPort := config.Envs.App.Port
		log.Info().Msgf("Server is running on port %s", serverPort)
		if err = e.Start(":" + serverPort); err != nil {
			log.Fatal().Msgf("Error while starting server: %v", err)
		}
	}()
	// End Run server in goroutine

	// Handle graceful shutdown
	quit := make(chan os.Signal, 1)

	shutdownSignals := []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGINT}
	if runtime.GOOS == "windows" {
		shutdownSignals = []os.Signal{os.Interrupt}
	}

	signal.Notify(quit, shutdownSignals...)
	<-quit
	log.Info().Msg("Server is shutting down ...")
	log.Info().Msg("Server gracefully stopped")
}
