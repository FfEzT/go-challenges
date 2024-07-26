package main

// TODO удалить тесты

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"

	"http-server/internal/config"
	"http-server/internal/http-server/handlers"
)

func main() {
  cfg := config.MustLoad()
  log := initLogger()

  router := chi.NewRouter()

  router.Use(middleware.RequestID)
  router.Use(middleware.Logger)
  router.Use(middleware.Recoverer)

  // NOTE ограничение числа запросов
  router.Use(
    httprate.Limit(
      cfg.Requests,             // requests
      time.Duration(cfg.Seconds) * time.Second, // per duration,
      httprate.WithLimitHandler(handlers.LimitHandler),
    ),
  )

  router.Post("/", handlers.CalculateHandler)

  // GraceFull ShutDown
  done := make(chan os.Signal, 1)
  signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

  srv := &http.Server{
    Addr:         cfg.Address,
    Handler:      router,
    ReadTimeout:  cfg.HTTP_server.Timeout,
    WriteTimeout: cfg.HTTP_server.Timeout,
    IdleTimeout:  cfg.HTTP_server.Idle_timeout,
  }

  go func() {
    if err := srv.ListenAndServe(); err != nil {
      log.Error("failed to start server")
    }
  }()

  log.Info("server started")

  <-done
  log.Info("stopping server")

  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()

  if err := srv.Shutdown(ctx); err != nil {
    log.Error("failed to stop server") // TODO error output

    return
  }
  log.Info("end")
}

func initLogger() *slog.Logger {
  return slog.New(
    slog.NewTextHandler(
      os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug},
    ),
  )
}