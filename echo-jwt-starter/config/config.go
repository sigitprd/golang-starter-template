package config

import (
	"echo-jwt-starter/pkg/config"
	"flag"
	"os"
	"strings"
	"sync"

	"github.com/rs/zerolog/log"
)

var (
	Envs *Config // Envs is global vars Config.
	once sync.Once
)

type Config struct {
	App struct {
		Name        string `env:"APP_NAME" required:"true"`
		Environment Env    `env:"APP_ENV" env-default:"production" required:"true"`
		BaseURL     string `env:"APP_BASE_URL" env-default:"http://localhost:3000" required:"true"`
		Port        string `env:"APP_PORT" required:"true"`
		LogLevel    string `env:"APP_LOG_LEVEL" env-default:"debug" required:"true"`
		LogFile     string `env:"APP_LOG_FILE" env-default:"./logs/app.log" required:"true"`
		BinDir      string `env:"APP_BIN_DIR" required:"true"`
	}
	APIKeys struct {
		XApiKey string `env:"X_API_KEY" required:"true"`
	}
	DB struct {
		Postgres struct {
			Host              string `env:"DB_HOST" env-default:"localhost" required:"true"`
			Port              string `env:"DB_PORT" env-default:"5432" required:"true"`
			Username          string `env:"DB_USER" env-default:"postgres" required:"true"`
			Password          string `env:"DB_PASS" env-default:"postgres" required:"true"`
			Database          string `env:"DB_NAME" env-default:"postgres" required:"true"`
			SslMode           string `env:"DB_SSL_MODE" env-default:"disable" required:"true"`
			ConnectionTimeout int    `env:"DB_CONN_TIMEOUT" env-default:"30" env-description:"database timeout in seconds" required:"true"`
			MaxOpenCons       int    `env:"DB_MAX_OPEN_CONS" env-default:"20" env-description:"database max open conn in seconds" required:"true"`
			MaxIdleCons       int    `env:"DB_MAX_IdLE_CONS" env-default:"20" env-description:"database max idle conn in seconds" required:"true"`
			ConnMaxLifetime   int    `env:"DB_CONN_MAX_LIFETIME" env-default:"0" env-description:"database conn max lifetime in seconds" required:"false"`
		}
	}
	Guard struct {
		JwtSecret         string `env:"JWT_SECRET" required:"true"`
		JwtTtlHours       int    `env:"JWT_TTL_HOURS" env-default:"24" required:"true"`        // 24 hours
		JwtRefreshTtlDays int    `env:"JWT_REFRESH_TTL_DAYS" env-default:"30" required:"true"` // 30 days
	}
}

// Option is Configure type return func.
type Option = func(c *Configure) error

// Configure is the data struct.
type Configure struct {
	path     string
	filename string
}

// Configuration create instance.
func Configuration(opts ...Option) *Configure {
	c := &Configure{}

	for _, opt := range opts {
		err := opt(c)
		if err != nil {
			panic(err)
		}
	}
	return c
}

// Initialize will create instance of Configure.
func (c *Configure) Initialize() {
	once.Do(func() {
		Envs = &Config{}
		if err := config.Load(config.Opts{
			Config:    Envs,
			Paths:     []string{c.path},
			Filenames: []string{c.filename},
		}); err != nil {
			log.Fatal().Err(err).Msg("get config error")
		}

		// Validate the loaded configuration
		if err := Envs.Validate(); err != nil {
			log.Fatal().Err(err).Msg("configuration validation error")
		}
	})
}

// WithPath will assign to field path Configure.
func WithPath(path string) Option {
	return func(c *Configure) error {
		c.path = path
		return nil
	}
}

// WithFilename will assign to field name Configure.
func WithFilename(name string) Option {
	return func(c *Configure) error {
		c.filename = name
		return nil
	}
}

// Load will load the configuration.
func LoadEnvs() (newArgs []string) {
	configPath := flag.String("config_path", "./", "path to config file")
	configFilename := flag.String("config_filename", ".env", "config file name")
	flag.Parse()

	var logCfg string
	if *configPath == "./" {
		logCfg = *configPath + *configFilename
	} else {
		logCfg = *configPath + "/" + *configFilename
	}

	log.Info().Msgf("Initializing configuration with config: %s", logCfg)

	Configuration(
		WithPath(*configPath),
		WithFilename(*configFilename),
	).Initialize()

	for _, arg := range os.Args {
		if strings.Contains(arg, "config_path") || strings.Contains(arg, "config_filename") {
			continue
		}

		newArgs = append(newArgs, arg)
	}

	return newArgs
}
