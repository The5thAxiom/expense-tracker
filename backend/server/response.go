package server

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	ErrorMessage *string `json:"errorMessage"`
	Data         any     `json:"data"`
}

func NewResponse(w http.ResponseWriter, statusCode int, data any, errorMessage *string) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")

	endpointResponse := Response{
		ErrorMessage: errorMessage,
		Data:         data,
	}

	jsonBytes, _ := json.Marshal(endpointResponse)
	w.Write(jsonBytes)

	// can the below be done in a middleware method??
	log.Printf("%d %s: %s", statusCode, http.StatusText(statusCode), string(jsonBytes))
}
