package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type EndpointResponse struct {
	ErrorMessage *string `json:"errorMessage"`
	Data         any     `json:"data"`
}

func NewEndpointResponse(w http.ResponseWriter, statusCode int, data any, errorMessage *string) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")

	endpointResponse := EndpointResponse{
		ErrorMessage: errorMessage,
		Data:         data,
	}

	jsonBytes, _ := json.Marshal(endpointResponse)
	w.Write(jsonBytes)

	// can the below be done in a middleware method??
	dataLogMessage := ""
	if data != nil {
		dataLogMessage = fmt.Sprintf("data: %s ", string(jsonBytes))
	}

	errorLogMessage := ""
	if errorMessage != nil {
		errorLogMessage = fmt.Sprintf("error: %s", *errorMessage)
	}
	log.Printf("%d %s: %s %s", statusCode, http.StatusText(statusCode), dataLogMessage, errorLogMessage)
}
