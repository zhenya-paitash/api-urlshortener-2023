package main

import (
	"log/slog"

	"github.com/zhenya-paitash/api-urlshortener-2023/internal/config"
	"github.com/zhenya-paitash/api-urlshortener-2023/internal/logger"
)

func main() {
	// NOTE: config: cleanenv
  config := config.MustLoad()

	// NOTE: logger: slog
  log := logger.SetupLogger(config.Env)
  log.Info("url-shortener initialized", slog.String("env", config.Env))
  log.Debug("debug messages are enabled")

	// TODO: storage
	// NOTE: logger: sqlite

	// TODO: router
	// NOTE: router: chi, chi-render

	// TODO: run server
}

