package handler

import (
    "log"
    "net/http"
)

func LogHandler(fn func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Printf("Received a request %s %s \n", r.Method, r.RequestURI)
        fn(w, r)
    }
}
