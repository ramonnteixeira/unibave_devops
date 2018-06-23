package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", healthCheckHandler)
	http.ListenAndServe(":5000", nil)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true, "teste": 1}`)
}
