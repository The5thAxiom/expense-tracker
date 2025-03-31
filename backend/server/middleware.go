package server

import (
	"log"
	"net/http"
)

func (s Server) UseMiddleware(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return s.LogRequest(handler)
}

func (s Server) LogRequest(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s request on %s\n", r.Method, r.URL)
		handler(w, r)
	}
}
