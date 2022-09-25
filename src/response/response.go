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
		"err":    "",
		"msg":    msg,
	}
	ctx.JSON(statusCode, res)
}

func Fail(ctx *gin.Context, statusCode int, errors []gin.H, msg string) {
	if ctx.Writer.Written() {
		log.Warn("response body was already written! will not overwrite")
		return
	}
	res := response{
		"status": "fail",
		"data": response{
			"errors": errors,
		},
		"msg": msg,
	}
	ctx.JSON(statusCode, res)
}
