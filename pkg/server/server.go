package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	port   string
	router *mux.Router
}

func NewServer(port string) Server {
	return Server{
		port:   fmt.Sprintf(":%s", port),
		router: mux.NewRouter().StrictSlash(true),
	}
}

func (s Server) Serve() {
	s.router.HandleFunc("/", homeLink)
	s.router.HandleFunc("/charts", chartLink)
	log.Fatal(http.ListenAndServe(s.port, s.router))
}
