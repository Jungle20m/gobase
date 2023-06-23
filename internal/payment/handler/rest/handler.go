package rest

import (
	"github.com/gin-gonic/gin"
	mutils "gobase/utils"
	"net/http"
)

func NewHandler(utils mutils.IUtils) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	handler := gin.New()

	user := handler.Group("/user")
	{
		user.GET("/", GetUser(utils))
	}

	handler.GET("/ping", func(c *gin.Context) { c.JSON(http.StatusOK, "pong") })

	return handler
}
