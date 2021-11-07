// Package middleware defines various pieces of middleware that conform
// to the http.Handler interface.
package middleware

import "net/http"

// Headers sets response headers that are important for the
// API to function properly.
func Headers(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}

func Options(w http.ResponseWriter, r *http.Request) {
	header := w.Header()

	// check for CORS headers
	if r.Header.Get("Access-Control-Request-Method") != "" {
		header.Set("Access-Control-Allow-Headers", "Content-Type")
		header.Set("Access-Control-Allow-Methods", "DELETE, GET, OPTIONS, POST")
		header.Set("Access-Control-Allow-Origin", "*")
		header.Set("Access-Control-Max-Age", "86400")
	}

	w.WriteHeader(http.StatusNoContent)
}
