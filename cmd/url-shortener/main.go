package main

import (
	"os"

	"github.com/zhenya-paitash/api-urlshortener-2023/internal/config"
	"github.com/zhenya-paitash/api-urlshortener-2023/internal/lib/logger/sl"
	"github.com/zhenya-paitash/api-urlshortener-2023/internal/logger"
	"github.com/zhenya-paitash/api-urlshortener-2023/internal/storage/sqlite"
)

func main() {
	// config: cleanenv
	config := config.MustLoad()

	// logger: slog
	log := logger.SetupLogger(config.Env)
	log.Info("url-shortener initialized", sl.Str("env", config.Env))
	log.Debug("debug messages are enabled")

	// storage: sqlite
	storage, err := sqlite.New(config.StoragePath)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}
  _ = storage

	// TODO: router
	// NOTE: router: chi, chi-render

	// TODO: run server
}
