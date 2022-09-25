package controller

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"zamp/request"
	"zamp/src/response"
	"zamp/src/service"
)

func GetUserCart(ctx *gin.Context) {
	userId := ctx.Query("user_id")

	results, err := service.GetUserCart(ctx, userId)
	if err != nil {
		response.Fail(ctx, http.StatusInternalServerError, err, "")
	}
	response.Success(ctx, http.StatusOK, results, "success")
}

func PostUserCart(ctx *gin.Context) {
	var postCartRequest request.PostCartRequest
	err := ctx.ShouldBindJSON(&postCartRequest)
	if err != nil {
		log.Errorln("invalid request")
		response.Fail(ctx, http.StatusBadRequest, err, "request validation failed")
		return
	}

	err = service.PostUserCart(ctx, postCartRequest)
	if err != nil {
		response.Fail(ctx, 500, err, "")
		return
	}
	response.Success(ctx, http.StatusOK, nil, "success")
}
