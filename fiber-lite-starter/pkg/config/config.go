package config

import (
	"os"
	"path/filepath"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Opts struct {
		Config    any
		Paths     []string
		Filenames []string
	}
)

func Load(opts Opts) error {
	for _, p := range opts.Paths {
		fp := filepath.Join(p, ".env")
		// load env from file
		if _, fileErr := os.Stat(fp); fileErr == nil {
			// Set ENV for development
			_ = cleanenv.ReadConfig(fp, opts.Config)
		}
	}
	var err error
	for _, f := range opts.Filenames {
		fileFound := false
		for _, p := range opts.Paths {
			fp := filepath.Join(p, f)
			if _, fileErr := os.Stat(fp); fileErr != nil {
				continue // Try next path
			}
			fileFound = true
			err = cleanenv.ReadConfig(fp, opts.Config)
			// Continue even if ReadConfig fails to allow env vars to override
		}
		if !fileFound {
			return os.ErrNotExist
		}
	}

	// Read environment variables to allow overriding config file values
	// This should always be attempted, even if config file reading failed
	// This allows environment variables to override config file values and
	// provides flexibility when config files have parsing errors
	if envErr := cleanenv.ReadEnv(opts.Config); envErr != nil && err == nil {
		// Only return env error if there wasn't already a config file error
		// If config file had errors, those take precedence in error reporting
		err = envErr
	}

	return err
}
