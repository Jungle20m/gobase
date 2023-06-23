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
	listener net.Listener
}

func NewInstance(host string, port int) *server {
	instance := grpc.NewServer()
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	return &server{
		instance: instance,
		listener: listener,
	}
}

func (s *server) GetInstance() *grpc.Server {
	return s.instance
}

func (s *server) Serve() {
	// Reflection fo automation generate url
	reflection.Register(s.instance)

	go func() {
		if err := s.instance.Serve(s.listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	log.Printf("Grpc server is listening in %v\n", s.listener.Addr())
}

func (s *server) Terminal() {
	s.instance.Stop()
	log.Println("Grpc server is closed")
}
