package response

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type response gin.H

func Success(ctx *gin.Context, statusCode int, data interface{}, msg string) {
	if ctx.Writer.Written() {
		log.Warn("response body was already written! will not overwrite")
		return
	}
	res := response{
		"status": "success",
		"data":   data,
		"msg":    msg,
	}
	ctx.JSON(statusCode, res)
}

func Fail(ctx *gin.Context, statusCode int, err error, msg string) {
	if ctx.Writer.Written() {
		log.Warn("response body was already written! will not overwrite")
		return
	}
	res := response{
		"status": "fail",
		"err":    err.Error(),
		"msg":    msg,
	}
	ctx.JSON(statusCode, res)
}
