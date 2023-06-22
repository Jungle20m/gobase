package rest

import (
	"fmt"
	"net/http"
	"time"
)

const (
	_defaultReadTimeout     = 10 * time.Second
	_defaultWriteTimeout    = 10 * time.Second
	_defaultAddr            = ":8001"
	_defaultShutdownTimeout = 5 * time.Second
)

type server struct {
	server *http.Server
	host   string
	port   int
}

func NewServer(handler http.Handler, host string, port int) *server {
	//handler := gin.New()
	//handler.GET("/ping", func(c *gin.Context) { c.JSON(http.StatusOK, "ping") })

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
		host:   host,
		port:   port,
	}
}

func (s *server) Serve() {
	go func() {
		s.server.ListenAndServe()
	}()

	fmt.Printf("Rest Server Listening In %v\n", s.server.Addr)
}

func (s *server) Terminal() {
	fmt.Println("Rest Server Stopped")
}
