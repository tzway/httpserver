package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type chirp struct {
	Body string `json:"body"`
}

type cleanedChirp struct {
	CleanedBody string `json:"cleaned_body"`
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
	var cp chirp
	err := decoder.Decode(&cp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respBody, _ := json.Marshal(struct {
			Error string `json:"error"`
		}{Error: "Invalid JSON"})
		w.Write(respBody)
		return
	}

	// replace bad words
	cleanedCp := cleanedChirp{CleanedBody: replaceBadWord(cp.Body)}

	w.WriteHeader(http.StatusOK)
	respBody, _ := json.Marshal(cleanedCp)
	w.Write(respBody)
}

func replaceBadWord(ori string) string {
	badWords := []string{"kerfuffle", "sharbert", "fornax"}
	sliceOri := strings.Split(ori, " ")
	for i, word := range sliceOri {
		for _, badWord := range badWords {
			if strings.ToLower(word) == badWord {
				sliceOri[i] = "****"
			}
		}
	}
	return strings.Join(sliceOri, " ")
}
