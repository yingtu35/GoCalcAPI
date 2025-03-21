package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"
)

type Logger struct {
	Next http.Handler // The next handler in the chain
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// record how long the request takes
	start := time.Now()
	l.Next.ServeHTTP(w, r)
	slog.Info(
		"incoming request",
		slog.String("method", r.Method),
		slog.String("path", r.URL.Path),
		slog.String("remote_addr", r.RemoteAddr),
		slog.Int64("time taken (ns)", time.Since(start).Nanoseconds()),
	)
}

func NewLogger(next http.Handler) *Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	slog.SetDefault(logger)
	return &Logger{Next: next}
}
