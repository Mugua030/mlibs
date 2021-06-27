package grpc

import (
	"fmt"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start() error {
	fmt.Println("[grpc.server] start...")
}

func (s *Server) Stop() error {
	fmt.Println("[grpc.server] stop...")
}
