package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zamp/src/response"
)

func GetProducts(ctx *gin.Context) {
	response.Success(ctx, http.StatusOK, nil, "success")
}
