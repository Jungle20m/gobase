package rest

import (
	"github.com/gin-gonic/gin"
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

func NewServer() *server {
	handler := gin.New()
	handler.GET("/ping", func(c *gin.Context) { c.JSON(http.StatusOK, "ping") })

	httpServer := &http.Server{
		Addr:              _defaultAddr,
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

}

func (s *server) Terminal() {

}
