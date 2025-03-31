package server

import (
	"encoding/json"
	"net/http"
)

type EndpointResponse struct {
	ErrorMessage *string `json:"errorMessage"`
	Data         any     `json:"data"`
}

func NewEndpointResponse(w http.ResponseWriter, statusCode int, data any, errorMessage *string) {
	w.WriteHeader(statusCode)

	endpointResponse := EndpointResponse{
		ErrorMessage: errorMessage,
		Data:         data,
	}

	jsonBytes, _ := json.Marshal(endpointResponse)
	w.Write(jsonBytes)
}
