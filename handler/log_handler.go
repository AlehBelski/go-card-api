package handler

import (
    "log"
    "net/http"
)

// LogHandler provides a wrapper to a general request handler
// with ability to log the incoming request.
func LogHandler(fn func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Printf("Received a request %s %s \n", r.Method, r.RequestURI)
        fn(w, r)
    }
}
