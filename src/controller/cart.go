package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zamp/src/response"
	"zamp/src/service"
)

func GetUserCart(ctx *gin.Context) {
	userId := ctx.Query("user_id")

	results, err := service.GetUserCart(ctx, userId)
	if err != nil {
		response.Fail(ctx, 500, err, "")
	}
	response.Success(ctx, 200, results, "")
	response.Success(ctx, http.StatusOK, nil, "success")
}
