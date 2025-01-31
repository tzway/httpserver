package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/tzway/httpserver/internal/database"
)

func main() {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var cfg *apiConfig = &apiConfig{
		dbQueries: database.New(db),
	}

	fileSever := http.FileServer(http.Dir("."))

	mux := http.NewServeMux()
	mux.Handle("/app/", http.StripPrefix("/app", cfg.middlewareMetricsInc(fileSever)))

	mux.HandleFunc("GET /admin/metrics", cfg.handleMetrics)
	mux.HandleFunc("POST /admin/reset", cfg.handleReset)

	mux.HandleFunc("GET /api/healthz", handleHealthz)

	mux.HandleFunc("POST /api/validate_chirp", handleValidateChirp)

	server := &http.Server{Handler: mux, Addr: ":8080"}
	log.Fatal(server.ListenAndServe())
}
