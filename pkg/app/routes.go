package app

import "net/http"

func (s *Server) RegisterRoutes() {
	http.HandleFunc("/", rootHandler())
}

func rootHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello you"))
	}
}
