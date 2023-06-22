package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewUserHandler() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	handler := gin.New()

	handler.GET("/ping", func(c *gin.Context) { c.JSON(http.StatusOK, "pong") })

	return handler
}
