package main

import (
	"log"
	"net/http"

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

	log.Fatal(http.ListenAndServe(":8080", rateLimiterMux))
}
