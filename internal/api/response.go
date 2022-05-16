package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func responseSuccess(ctx *gin.Context, msg string, data gin.H) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"data":    data,
		"message": msg,
	})
}

func responseFail(ctx *gin.Context, msg string, data gin.H) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    400,
		"data":    data,
		"message": msg,
	})
}
