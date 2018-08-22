package main

import (
	"net/http"

	"github.com/JonathonGore/httputil"
)

func main() {
	http.HandleFunc("/success", func(w http.ResponseWriter, r *http.Request) {
		httputil.Success(w)
	})

	http.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		httputil.Error(w, "unable to complete operation", 500)
	})

	http.ListenAndServe(":3000", nil)
}
