package config

import (
	"errors"
	"os"
)

const defaultHTTPAddr = ":8080"

var ErrNoDatabaseURL = errors.New("DATABASE_URL is not set")

type Config struct {
	HTTPAddr    string
	DatabaseURL string
}

func Load() (Config, error) {
	cfg := Config{
		HTTPAddr:    os.Getenv("HTTP_ADDR"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}

	if cfg.HTTPAddr == "" {
		cfg.HTTPAddr = defaultHTTPAddr
	}

	if cfg.DatabaseURL == "" {
		return Config{}, ErrNoDatabaseURL
	}

	return cfg, nil
}
