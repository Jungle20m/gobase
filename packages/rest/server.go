package rest

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	_defaultReadTimeout     = 10 * time.Second
	_defaultWriteTimeout    = 10 * time.Second
	_defaultAddr            = ":8000"
	_defaultShutdownTimeout = 5 * time.Second
)

type server struct {
	server *http.Server
}

func NewInstance(handler http.Handler, host string, port int) *server {
	httpServer := &http.Server{
		Addr:              fmt.Sprintf("%s:%d", host, port),
		Handler:           handler,
		TLSConfig:         nil,
		ReadTimeout:       _defaultReadTimeout,
		ReadHeaderTimeout: 0,
		WriteTimeout:      _defaultWriteTimeout,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}

	return &server{
		server: httpServer,
	}
}

func (s *server) Serve() {
	go func() {
		if err := s.server.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				log.Println("Rest server is closed")
			} else {
				log.Fatalf("Failed to run rest server: %s\n", err)
			}
		}
	}()
	log.Printf("Rest server is listening in %v\n", s.server.Addr)
}

func (s *server) Terminal() {
	s.server.Close()
}
