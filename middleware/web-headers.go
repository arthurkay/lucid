package middleware

import (
	"net/http"
)

// DefaultHeaders sets the application default headers set for app idenitification
func DefaultWebHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("User-Agent", "Lucid API Framework")
		w.Header().Set("X-Powered-By", "GO; golang 1.16")
		next.ServeHTTP(w, r)
	})
}
