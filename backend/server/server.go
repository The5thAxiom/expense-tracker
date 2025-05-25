package server

import (
	"backend/db"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/MicahParks/keyfunc"
)

type Handler func(w http.ResponseWriter, r *http.Request)
type MiddleWare func(handler Handler) Handler

type Server struct {
	Db db.DB

	Host string
	Port int

	mux        *http.ServeMux
	middleware []MiddleWare

	jwks     *keyfunc.JWKS
	audience string
}

type ServerOptions struct {
	StaticDir *string
	StaticUrl *string
}

func New(host string, port int, db db.DB, options *ServerOptions) Server {
	jwksUrl := os.Getenv("AUTH0_PUBKEY_URL")
	jwks, err := keyfunc.Get(jwksUrl, keyfunc.Options{
		RefreshInterval: time.Hour,
		RefreshErrorHandler: func(err error) {
			fmt.Printf("JWKS refresh, error: %s\n", err)
		},
		RefreshUnknownKID: true,
	})
	if err != nil {
		log.Fatalf("Error occurred loading public key: %v", err)
	}

	server := Server{
		Db: db,

		Host: host,
		Port: port,

		mux:        http.NewServeMux(),
		middleware: make([]MiddleWare, 0),

		jwks:     jwks,
		audience: os.Getenv("AUTH0_AUDIENCE"),
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

	err := http.ListenAndServe(s.Address(), s.mux)
	if err != nil {
		log.Fatal(err)
	}
}

func (s Server) addRouteWithMiddleware(route string, middleware []MiddleWare, handler Handler) {
	finalHandler := handler
	for _, m := range middleware {
		finalHandler = m(finalHandler)
	}
	for _, m := range s.middleware {
		finalHandler = m(finalHandler)
	}
	s.mux.HandleFunc(route, finalHandler)
}

func (s Server) addRoute(route string, handler Handler) {
	finalHandler := handler
	for _, m := range s.middleware {
		finalHandler = m(finalHandler)
	}
	s.mux.HandleFunc(route, finalHandler)
}

func (s *Server) useMiddleware(m MiddleWare) {
	s.middleware = append(s.middleware, m)
}
