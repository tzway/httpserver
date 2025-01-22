package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir("."))))

	mux.HandleFunc("/healthz", handler)

	server := &http.Server{Handler: mux, Addr: ":8080"}
	server.ListenAndServe()
}

// handler is a simple http handler that returns a 200 status code without no requirement of the request
func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
