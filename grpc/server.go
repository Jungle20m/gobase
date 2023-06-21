package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type server struct {
	instance *grpc.Server
	host     string
	port     int
}

func NewServer() *server {
	instance := grpc.NewServer()

	return &server{
		instance: instance,
		host:     "",
		port:     9001,
	}
}

func (s *server) Instance() *grpc.Server {
	return s.instance
}

func (s *server) Serve() {

	// Reflection fo automation generate url
	reflection.Register(s.instance)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := s.instance.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) Terminal() {

}
