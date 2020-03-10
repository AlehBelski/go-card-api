// Package middleware provides the middleware functions that add
// additional behavior over basic handle functions.
package middleware

import (
	"log"
	"net/http"
)

// LogMiddleware provides a middleware to a general request handler
// with ability to log the incoming request.
func LogMiddleware(fn func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received a request %s %s \n", r.Method, r.RequestURI)
		fn(w, r)
	}
}
