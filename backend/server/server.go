package server

import (
	"backend/database"
	"fmt"
	"log"
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

func New(port int, db database.DB, options *ServerOptions) Server {
	server := Server{
		Db:   db,
		Mux:  http.NewServeMux(),
		Port: port,
	}

	// fs := http.FileServer(http.Dir(""))

	server.AddEndpoints()

	return server
}

func (s Server) Run() {
	hostAddress := fmt.Sprintf(":%d", s.Port)
	fmt.Printf("serving Y3VudA== on port %d...\n", s.Port)

	err := http.ListenAndServe(hostAddress, s.Mux)
	if err != nil {
		log.Fatal(err)
	}
}
