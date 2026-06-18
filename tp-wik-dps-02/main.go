package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		notFound(w)
		return
	}

	headers := make(map[string][]string)

	for k, v := range r.Header {
		headers[k] = v
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(headers)
}

func notFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
}

func main() {
	port := os.Getenv("PING_LISTEN_PORT")
	if port == "" {
		port = "9090"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/ping", pingHandler)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	fmt.Printf("Server running at http://localhost:%s/ping\n", port)

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
