package main

import (
	"log/slog"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/zhenya-paitash/api-urlshortener-2023/internal/config"
	"github.com/zhenya-paitash/api-urlshortener-2023/internal/http-server/handlers/url/save"
	mwHTTPLogger "github.com/zhenya-paitash/api-urlshortener-2023/internal/http-server/middleware"
	"github.com/zhenya-paitash/api-urlshortener-2023/internal/lib/logger/sl"
	"github.com/zhenya-paitash/api-urlshortener-2023/internal/logger"
	"github.com/zhenya-paitash/api-urlshortener-2023/internal/storage/sqlite"
)

func main() {
	// config: cleanenv
	config := config.MustLoad()

	// logger: slog
	log := logger.SetupLogger(config.Env)
	log.Info("url-shortener initialized", slog.String("env", config.Env))
	log.Debug("debug messages are enabled")

	// storage: sqlite
	storage, err := sqlite.New(config.StoragePath)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}
	_ = storage

	// router: chi, chi-render
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	// router.Use(middleware.Logger)
	router.Use(mwHTTPLogger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Post("/url", save.New(log, storage))

	// TODO: run server
}
