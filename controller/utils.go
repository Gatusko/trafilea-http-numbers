package controller

import (
	"encoding/json"
	"github.com/Gatusko/trafilea-http-numbers/domain/model"
	"log"
	"net/http"
)

// We respond always with a json body
func respondWithJson(w http.ResponseWriter, code int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshaling json %v", err)
	}
	w.WriteHeader(code)
	w.Write(data)
}

// We respond always with a json body
func respondWithoutBody(w http.ResponseWriter, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
}

// Decode this will disallow unkown fields
func readJson(r *http.Request, data any) error {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(data)
	if err != nil {
		return err
	}
	return err
}

// We send an error message and then respond with respond with json
func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Server error 5xx : %v", code)
	}
	res := model.NewError(msg)
	respondWithJson(w, code, res)
}
