package rest

import (
	"github.com/gin-gonic/gin"
	"gobase/internal/payment/repository"
	"gobase/internal/payment/usecase"
	mutils "gobase/utils"
	"net/http"
)

func GetUser(utils mutils.IUtils) gin.HandlerFunc {
	return func(c *gin.Context) {
		//var body registerBody
		//if err := c.ShouldBind(&body); err != nil {
		//	c.JSON(http.StatusBadRequest, httpsdk.HttpErrorResponse(err))
		//	return
		//}

		repo := repository.NewInstance(repository.WithGorm(utils.GetDB()))
		uc := usecase.NewUserUC(utils, repo)

		uc.GetUserByID(c.Request.Context(), 1)

		c.JSON(http.StatusOK, "ok")
	}
}
