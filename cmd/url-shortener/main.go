package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/zhenya-paitash/api-urlshortener-2023/internal/config"
	"github.com/zhenya-paitash/api-urlshortener-2023/internal/http-server/handlers/redirect"
	"github.com/zhenya-paitash/api-urlshortener-2023/internal/http-server/handlers/url/save"
	deleteHandler "github.com/zhenya-paitash/api-urlshortener-2023/internal/http-server/handlers/url/delete"
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

	// router: chi, chi-render
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	// router.Use(middleware.Logger)
	router.Use(mwHTTPLogger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Post("/url", save.New(log, storage))
	router.Get("/{alias}", redirect.New(log, storage))
  router.Delete("/url/{alias}", deleteHandler.New(log, storage))

	// server
	log.Info("starting server", slog.String("address", config.Address))
	srv := &http.Server{
		Addr:         config.Address,
		Handler:      router,
		ReadTimeout:  config.HTTPServer.Timeout,
		WriteTimeout: config.HTTPServer.Timeout,
		IdleTimeout:  config.HTTPServer.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server")
	}

	log.Error("server stoped")
}
