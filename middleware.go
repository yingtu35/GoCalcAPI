package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"
)

type ResponseWriterWithStatus struct {
	http.ResponseWriter     // Embed the http.ResponseWriter
	statusCode          int // The status code of the response
}

func (w *ResponseWriterWithStatus) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

type Logger struct {
	Next http.Handler // The next handler in the chain
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Wrap the response writer with our custom implementation
	ws := &ResponseWriterWithStatus{ResponseWriter: w, statusCode: http.StatusOK}
	// record how long the request takes
	start := time.Now()
	l.Next.ServeHTTP(ws, r)
	slog.Info(
		"incoming request",
		slog.String("method", r.Method),
		slog.String("path", r.URL.Path),
		slog.String("remote_addr", r.RemoteAddr),
		slog.Int64("time taken (ns)", time.Since(start).Nanoseconds()),
		slog.Int("status_code", ws.statusCode),
	)
}

func NewLogger(next http.Handler) *Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	slog.SetDefault(logger)
	return &Logger{Next: next}
}
