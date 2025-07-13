package main

import (
	"fiber-lite-starter/config"
	"fiber-lite-starter/internal/repository/psql"
	"fiber-lite-starter/internal/routes"
	"fiber-lite-starter/middleware"
	dbconfig "fiber-lite-starter/pkg/db"
	"fiber-lite-starter/pkg/logging"
	"fiber-lite-starter/pkg/validator"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
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

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName: config.Envs.App.Name,
	})

	// Middleware
	// Application Middlewares
	if config.Envs.App.Environment.IsProd() {
		app.Use(limiter.New(limiter.Config{
			Max:        50,
			Expiration: 30 * time.Second,
		}))
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,PATCH,OPTIONS,HEAD",
		AllowHeaders: "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin,Authorization,x-api-key",
	}))
	app.Use(middleware.ValidatorMiddleware(validator.NewValidator()))
	app.Use(compress.New())
	app.Use(requestid.New())
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
			log.Error().Interface("error", e).Msg("Panic occurred")
		},
	}))
	//if config.Envs.App.Environment.IsLocal() {
	//	app.Use(logger.New())
	//}

	// Register routes
	routeRegistry := routes.NewRouteRegistry(
		psql.NewRepositoryRegistry(db),
	)
	routeRegistry.RegisterRoutes(app)

	// Run server in goroutine
	go func() {
		port := config.Envs.App.Port
		log.Info().Msgf("Server is running on port %s", port)
		if err := app.Listen(":" + port); err != nil {
			log.Fatal().Err(err).Msg("Error while starting server")
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	shutdownSignals := []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGINT}
	if runtime.GOOS == "windows" {
		shutdownSignals = []os.Signal{os.Interrupt}
	}
	signal.Notify(quit, shutdownSignals...)
	<-quit

	log.Info().Msg("Server is shutting down...")
	log.Info().Msg("Server gracefully stopped")
}
