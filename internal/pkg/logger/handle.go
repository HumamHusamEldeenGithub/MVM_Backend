package logger

import (
	"log"
	"net/http"
	"time"
)

func ServerLogger() http.HandlerFunc {
	// Create a logger middleware that logs incoming requests and their responses
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the incoming request
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)

		// Call the next handler in the chain
		startTime := time.Now()
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.DefaultServeMux.ServeHTTP(w, r)
		})
		next.ServeHTTP(w, r)

		// Log the response
		duration := time.Now().Sub(startTime)
		log.Printf("%s %s %s %s %s", r.RemoteAddr, r.Method, r.URL, w.Header().Get("status"), duration)
	})
}
