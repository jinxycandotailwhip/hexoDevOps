package server

import (
	"hexoDevOps/handlers"
	"log"
	"net/http"
)

type Server struct{}

func (s *Server) StartServe() {
	log.Println("start server")
	http.HandleFunc("/api/regenerate", handlers.Regenerate)
	http.HandleFunc("/api/new_page", handlers.NewPage)
	err := http.ListenAndServe("0.0.0.0:4001", nil)
	if err != nil {
		log.Fatal("listen and server", err)
	}
	log.Println("server closed")
}
