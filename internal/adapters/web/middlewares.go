package web

import (
	"log"
	"net/http"
	"time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func CommonMiddleware() Middleware {
	return func(nextMiddleware http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "application/json")
			nextMiddleware(w, r)
		}
	}
}

func LoggingMiddleware() Middleware {
	return func(nextMiddleware http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() { log.Println(r.URL.Path, time.Since(start).Round(time.Millisecond)) }()
			nextMiddleware(w, r)
		}
	}
}
