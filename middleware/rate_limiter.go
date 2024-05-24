package middleware

import (
	"net/http"
	"github.com/aureliomalheiros/rate-limiter/limiter"
)

func RateLimiterMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rl := limiter.NewRateLimiter()
		ip := r.RemoteAddr
		token := r.Header.Get("API_KEY")

		if !rl.Allow(ip, token) {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("you have reached the maximum number of requests or actions allowed within a certain time frame"))
			return
		}

		next.ServeHTTP(w, r)
	})
}
