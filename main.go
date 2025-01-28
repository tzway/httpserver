package main

import (
	"log"
	"net/http"
)

func main() {
	fileSever := http.FileServer(http.Dir("."))
	var cfg *apiConfig = &apiConfig{}
	mux := http.NewServeMux()
	mux.Handle("/app/", http.StripPrefix("/app", cfg.middlewareMetricsInc(fileSever)))

	mux.HandleFunc("GET /admin/metrics", cfg.handleMetrics)
	mux.HandleFunc("POST /admin/reset", cfg.handleReset)

	mux.HandleFunc("GET /api/healthz", handleHealthz)

	server := &http.Server{Handler: mux, Addr: ":8080"}
	log.Fatal(server.ListenAndServe())
}
