package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zamp/src/response"
	"zamp/src/service"
)

func GetProductDetails(ctx *gin.Context) {
	productId := ctx.Param("productId")

	results, err := service.GetProductDetails(ctx, productId)
	if err != nil {
		response.Fail(ctx, 500, err, "")
	}
	response.Success(ctx, 200, results, "")
	response.Success(ctx, http.StatusOK, nil, "success")
}
