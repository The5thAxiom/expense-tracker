package server

import (
	"backend/database"
	"encoding/json"
	"net/http"
)

type Server struct {
	Db   database.DB
	Mux  *http.ServeMux
	Port int
}

type ServerOptions struct {
	StaticDir *string
	StaticUrl *string
}

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
