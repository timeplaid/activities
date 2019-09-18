package app

import (
	"fmt"
	"net/http"
	"strconv"
)

type Server struct {
	router http.ServeMux
	port   int
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) ListenAndServe() {
	var err = http.ListenAndServe(fmt.Sprintf(":%s", strconv.Itoa(s.port)), s)
	if err != nil {
		fmt.Println("unable to launch http server due to error", err)
	}
}

func NewServer() Server {
	return Server{port: 8080}
}
