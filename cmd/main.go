package main

import (
	"context"
	"errors"
	"github.com/EDEN-NN/hydra-api/internal/di"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	container, err := di.NewContainer(ctx)
	if err != nil {
		log.Fatalf("API fail to start: %v", err)
	}

	srv := &http.Server{
		Addr:         container.AppConfig.ServerPort,
		Handler:      container.Router,
		ReadTimeout:  container.AppConfig.ReadTimeout,
		WriteTimeout: container.AppConfig.WriteTimeout,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server down: %v", err)
		}
	}()

	<-ctx.Done()
	log.Fatalf("Shutting down...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Printf("Shutdown error: %v", err)
	} else {
		log.Println("Server shutting down correctly")
	}
}
