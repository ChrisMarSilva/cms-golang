package controllers

import (
	"encoding/json"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
	// RespondWithJSON(w, code, struct{Erro string ´json:"erro"{Erro: message}´})
}

func RespondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(statusCode)
	if payload != nil{
		json.NewEncoder(w).Encode(payload)
	}
	// response, _ := json.Marshal(payload)
	// w.Write(response)
}
