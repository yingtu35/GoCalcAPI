package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yingtu35/GoCalcAPI/internal/api"
	"github.com/yingtu35/GoCalcAPI/internal/middleware"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", api.HealthCheckHandler)
	mux.HandleFunc("/add", api.AddHandler)
	mux.HandleFunc("/subtract", api.SubtractHandler)
	mux.HandleFunc("/multiply", api.MultiplyHandler)
	mux.HandleFunc("/divide", api.DivideHandler)

	LoggerMux := middleware.NewLogger(mux)
	rateLimiterMux := middleware.NewRateLimiter(LoggerMux, 10, 5)

	server := &http.Server{
		Addr:    ":8080",
		Handler: rateLimiterMux,
	}

	// Graceful shutdown
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Start the server in a goroutine so that it doesn't block
	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	// Wait for the signal
	<-sigs

	// Create a context with a timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown the server
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Shutdown(): %v", err)
	}

	// Stop the rate limiter
	rateLimiterMux.Stop()
	log.Println("Server shutdown successfully")
}
