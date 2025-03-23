package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

type RateLimiter struct {
	Next          http.Handler   // The next handler in the chain
	Limit         int            // The number of requests allowed in the duration
	BurstyLimit   int            // The number of bursty requests allowed
	BurstyLimiter chan time.Time // The channel to manage the bursty requests
	Ticker        *time.Ticker   // The ticker to add tokens
}

func (r *RateLimiter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// if burstyLimiter is empty, then the request is rejected
	select {
	case <-r.BurstyLimiter:
		r.Next.ServeHTTP(w, req)
		return
	default:
		http.Error(w, "Too many requests", http.StatusTooManyRequests)
		slog.Error("Too many requests",
			slog.String("method", req.Method),
			slog.String("path", req.URL.Path),
			slog.String("remote_addr", req.RemoteAddr),
			slog.Int("status_code", http.StatusTooManyRequests),
		)
		return
	}
}

func NewRateLimiter(next http.Handler, limit int, burstyLimit int) *RateLimiter {

	burstyLimiter := make(chan time.Time, burstyLimit)

	// Create a ticker that adds one token per (1000/limit) milliseconds
	tickInterval := time.Second / time.Duration(limit)
	ticker := time.NewTicker(tickInterval)

	// create a goroutine to manage the burstyLimiter
	go func() {
		for t := range ticker.C {
			select {
			case burstyLimiter <- t:
			default:
			}
		}
	}()

	return &RateLimiter{
		Next:          next,
		Limit:         limit,
		BurstyLimit:   burstyLimit,
		BurstyLimiter: burstyLimiter,
		Ticker:        ticker,
	}
}
