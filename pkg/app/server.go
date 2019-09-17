package app

import (
	"fmt"
	"net/http"
)

// chared dependencies
type Server struct {
	port string
}

func (s *Server) ListenAndServe() {
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		fmt.Println("unable to launch http server due to error", err)
	}
}

func NewServer(port string) Server {
	return Server{port: port}
}
