package logger

import (
	"log"
	"net/http"
	"time"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the incoming request
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)

		// Call the next handler in the chain
		startTime := time.Now()

		next.ServeHTTP(w, r)

		// Log the response
		duration := time.Now().Sub(startTime)
		log.Printf("%s %s %s %s %s", r.RemoteAddr, r.Method, r.URL, w.Header().Get("status"), duration)
	})
}
