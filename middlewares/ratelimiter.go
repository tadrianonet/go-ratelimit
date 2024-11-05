package middlewares

import (
	"fmt"
	"net/http"

	"golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(10, 1) // 10 requisições por segundo, burst de 1

func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			w.WriteHeader(http.StatusTooManyRequests)
			fmt.Fprintf(w, "429 - Too Many Requests")
			return
		}
		next.ServeHTTP(w, r)
	})
}
