package main

import (
	"encoding/json"
	"net/http"
)

type chirp struct {
	Body string `json:"body"`
}

func handleValidateChirp(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	// validate the chirp

	// if the content-length is longer than 140 characters, return a 400 status code
	if r.ContentLength > 140 {
		w.WriteHeader(http.StatusBadRequest)
		respBody, _ := json.Marshal(struct {
			Error string `json:"error"`
		}{Error: "Chirp is too long"})
		w.Write(respBody)
		return
	}

	// if the chirp is valid, return a 200 status code
	// if the chirp is invalid, return a 400 status code
	decoder := json.NewDecoder(r.Body)
	var c chirp
	err := decoder.Decode(&c)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respBody, _ := json.Marshal(struct {
			Error string `json:"error"`
		}{Error: "Invalid JSON"})
		w.Write(respBody)
		return
	}

	w.WriteHeader(http.StatusOK)
	respBody, _ := json.Marshal(struct {
		Valid bool `json:"valid"`
	}{Valid: true})
	w.Write(respBody)
}
