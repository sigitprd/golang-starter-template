package config

import (
	"database/sql"
	"fiber-jwt-starter/config"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

func NewPostgresConnection() (*sql.DB, error) {
	cfg := config.Envs
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DB.Postgres.Host,
		cfg.DB.Postgres.Port, // Ganti ke %s kalau Port bertipe string
		cfg.DB.Postgres.Username,
		cfg.DB.Postgres.Password,
		cfg.DB.Postgres.Database,
		cfg.DB.Postgres.SslMode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.DB.Postgres.MaxOpenCons)
	db.SetMaxIdleConns(cfg.DB.Postgres.MaxIdleCons)
	db.SetConnMaxLifetime(time.Duration(cfg.DB.Postgres.ConnMaxLifetime) * time.Minute)

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
