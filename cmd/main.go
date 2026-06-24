package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ankur59/students-api/internal/config"
)

func main() {
	cfg := config.MustLoad()

	router := http.NewServeMux()
	done := make(chan os.Signal, 1)

	// Learn about this signals
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}
	go func() {

		err := server.ListenAndServe()

		if err != nil {
			log.Fatal("Server failed to start")
		}
	}()

	<-done

	slog.Info("Gracefully shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Failed to shut  down server", slog.String("error", err.Error()))
	}

	slog.Info("server shutdown successfully")
}
