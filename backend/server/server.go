package server

import (
	"backend/db"
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	Db   db.DB
	Mux  *http.ServeMux
	Host string
	Port int
}

type ServerOptions struct {
	StaticDir *string
	StaticUrl *string
}

func New(host string, port int, db db.DB, options *ServerOptions) Server {
	server := Server{
		Db:   db,
		Mux:  http.NewServeMux(),
		Host: host,
		Port: port,
	}

	// fs := http.FileServer(http.Dir(""))

	server.AddEndpoints()

	return server
}

func (s Server) Address() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

func (s Server) Run() {
	fmt.Printf("serving Y3VudA== on host %s port %d...\n", s.Host, s.Port)

	err := http.ListenAndServe(s.Address(), s.Mux)
	if err != nil {
		log.Fatal(err)
	}
}
