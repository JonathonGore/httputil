package httputil

import (
	"log"
	"net/http"
	"time"
)

// Log logs the results and timing of the HTTP request to the configured logger.
func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lrw := newLoggingResponseWriter(w)

		start := time.Now()
		handler.ServeHTTP(lrw, r)
		statusCode := lrw.statusCode

		query := ""
		if r.URL.RawQuery != "" {
			query = "?" + r.URL.RawQuery
		}

		log.Printf("%v %v %v%v took %v seconds", statusCode, r.Method, r.URL.Path, query, time.Since(start).Seconds())
	})
}

// JSONResponse adds a JSON Content-Type header to wrapped requests.
func JSONResponse(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		handler.ServeHTTP(w, r)
	})
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

