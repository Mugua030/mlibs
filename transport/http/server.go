package http

import (
	"fmt"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start() error {
	fmt.Println("[http.server] start...")
}

func (s *Server) Stop() error {
	fmt.Println("[http.server] stop...")
}
